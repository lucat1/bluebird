package chess

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"github.com/notnil/chess"
	chessimage "github.com/notnil/chess/image"
	"github.com/rogpeppe/misc/svg"
)

type delayFunc func()

func (m *Match) delay(fn delayFunc) {
	if m.ticking.Load() {
		m.cancel <- true
	}
	m.cancel = make(chan bool)

	m.ticking.Store(true)
	go func() {
		select {
		case <-m.cancel:
			m.ticking.Store(false)

		case <-time.After(m.EndsAt.Sub(time.Now())):
			m.ticking.Store(false)
			fn()
		}
	}()
}

func (m *Match) periodic() {
	go func() {
		for {
			select {
			case <-m.quit:
				return

			case <-time.After(time.Second * time.Duration(ReplyPollInterval)):
				if err := m.FetchTweets(); err != nil {
					log.Printf("Faield to fetch replies to tweet %s: %v", m.TweetID, err)
				}
				break
			}
		}
	}()
}

func (m *Match) move(move string) error {
	if err := m.Game.MoveStr(move); err != nil {
		return err
	}

	log.Printf("Moved %s", move)
	if m.Game.Outcome() == chess.NoOutcome {
		m.EndsAt = time.Now().UTC().Add(m.Duration)
		m.Tweets = []request.Tweet{}
		m.Moves = map[string]uint{}
	}
	m.sendUpdate()
	m.PostGame()
	if m.Game.Outcome() == chess.NoOutcome {
		m.delay(m.onTurnEnd)
	} else {
		m.end()
	}
	return nil
}

func (m *Match) PlayerMove(move string) error {
	if m.Game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	return m.move(move)
}

func moveFromText(text string) *string {
	parts := strings.Split(text, " ")
	if len(parts) < 2 {
		return nil
	}
	return &parts[1]
}

func (m *Match) FetchTweets() (err error) {
	log.Println("Fetching chess replies")
	twts, err := request.Replies(m.TweetID, math.MaxInt, nil, nil)
	if err != nil {
		return
	}
	m.Tweets = []request.Tweet{}
	m.Moves = map[string]uint{}
	for _, twt := range twts {
		clone := m.Game.Clone()
		first := moveFromText(twt.Text)
		if first == nil {
			continue
		}
		if err := clone.MoveStr(*first); err == nil {
			tw := twt
			tw.Text = *first
			m.Tweets = append(m.Tweets, tw)
			m.Moves[*first]++
		}
	}
	m.sendUpdate()
	return
}

func (m *Match) randomMove() *string {
	log.Printf("No move given, playing random")
	moves := m.Game.ValidMoves()
	if len(moves) == 0 {
		return nil
	}
	randMove := moves[rand.Intn(len(moves))]
	if randMove == nil {
		return nil
	}
	move := chess.AlgebraicNotation{}.Encode(m.Game.Position(), randMove)
	return &move
}

func (m *Match) PostGame() {
	var msg string
	if m.Game.Outcome() == chess.NoOutcome {
		if m.Game.Position().Turn() == playerColor {
			msg = "Il pubblico ha scelto!"
		} else {
			msg = "Il giocatore ha fatto la sua mossa, ora tocca al popolo!"
		}
	} else {
		if m.Game.Position().Turn() == playerColor || m.Forfeited {
			msg = "Il pubblico ha vinto!"
			if m.Forfeited {
				msg += " (forfeit del giocatore)"
			}
		} else {
			msg = "Il giocatore ha vinto!"
		}
	}
	image, err := m.Image()
	if err != nil {
		log.Printf("WARN: Could not generate chessboard picture: %v", err)
		return
	}

	media, err := request.PostMedia(image)
	if err != nil {
		log.Printf("WARN: Could not upload media : %v", err)
		return
	}
	tweet, err := request.PostCustom(request.TweetRequest{
		Text: msg,
		Media: request.TweetRequestMedia{
			MediaIDs: []string{media.MediaID},
		},
	})
	if err != nil {
		log.Printf("WARN: Could not post a new tweet: %v", err)
		return
	}
	m.TweetID = tweet.ID
}

const (
	imageWidth  = 2048
	imageHeight = 2048
)

/*
export CGO_CFLAGS_ALLOW=".*"
export CGO_LDFLAGS_ALLOW=".*"
*/
func (m *Match) Image() (buf []byte, err error) {
	in, out := bytes.NewBuffer(buf), new(bytes.Buffer)
	if err = chessimage.SVG(in, m.Game.Position().Board()); err != nil {
		return
	}
	size := image.Point{imageWidth, imageHeight}
	dest, _ := svg.Render(in, size)
	// Create Writer from file
	b := bufio.NewWriter(out)
	// Write the image into the buffer
	err = png.Encode(b, dest)
	if err != nil {
		return
	}
	err = b.Flush()
	if err != nil {
		return
	}
	return out.Bytes(), nil
}

func (m *Match) sendUpdate() {
	select {
	case m.updates <- true:
	default:
	}
}

func (m *Match) Update() bool {
	return <-m.updates
}
