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
	"sync/atomic"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"github.com/notnil/chess"
	chessimage "github.com/notnil/chess/image"
	"github.com/rogpeppe/misc/svg"
)

const (
	playerColor       = chess.White
	stateFile         = "chess.json"
	ReplyPollInterval = 15 // seconds
)

var match *Match = nil

type Match struct {
	Code     string
	Duration time.Duration
	EndsAt   time.Time
	Game     *chess.Game
	TweetID  string
	Tweets   []request.Tweet

	timeout chan bool
	ticking atomic.Bool

	updates chan bool
	quit    chan bool
}

func (m *Match) delay() {
	if m.ticking.Load() {
		m.timeout <- true
	}
	m.timeout = make(chan bool)

	m.ticking.Store(true)
	go func() {
		select {
		case <-m.timeout:
			m.ticking.Store(false)

		case <-time.After(m.EndsAt.Sub(time.Now())):
			m.ticking.Store(false)
			m.onTurnEnd()
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
	for _, twt := range twts {
		clone := m.Game.Clone()
		first := moveFromText(twt.Text)
		if first == nil {
			continue
		}
		if err := clone.MoveStr(*first); err == nil {
			tw := twt
			first := moveFromText(twt.Text)
			if first == nil {
				continue
			}
			tw.Text = *first
			m.Tweets = append(m.Tweets, tw)
		}
	}
	m.sendUpdate()
	return
}

func (m *Match) getMoves() (moves map[string]uint, err error) {
	moves = map[string]uint{}
	tweets, err := request.Replies(m.TweetID, math.MaxInt, nil, nil)
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		clone := m.Game.Clone()
		first := moveFromText(tweet.Text)
		if first == nil {
			continue
		}
		if err := clone.MoveStr(*first); err == nil {
			moves[*first]++
		}
	}
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

func (m *Match) onTurnEnd() {
	var move *string
	if m.Game.Position().Turn() == playerColor {
		// Play a random move if the palyer didn't pick in time
		move = m.randomMove()
	} else {
		moves, err := m.getMoves()
		if err != nil {
			log.Printf("Could not get tweets replies: %v", err)
			return
		}
		if len(moves) != 0 {
			var mostValued uint
			for mv, val := range moves {
				if val > mostValued {
					move = &mv
					mostValued = val
				}
			}
		} else {
			move = m.randomMove()
		}
	}
	if move != nil {
		if err := m.Move(*move); err != nil {
			log.Printf("WARN: move %s failed:  %v", *move, err)
		}
	} else {
		// should never be reached
		m.close()
	}
}

func (m *Match) Move(move string) error {
	if err := m.Game.MoveStr(move); err != nil {
		return err
	}

	log.Printf("Moved %s", move)
	if m.Game.Outcome() == chess.NoOutcome {
		m.EndsAt = time.Now().UTC().Add(m.Duration)
		m.Tweets = []request.Tweet{}
	}
	m.sendUpdate()
	m.PostGame()
	if m.Game.Outcome() == chess.NoOutcome {
		m.delay()
	} else {
		m.close()
	}
	return nil
}

func (m *Match) PlayerMove(move string) error {
	if m.Game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	return m.Move(move)
}

func (m *Match) PostGame() {
	var msg string
	status := m.Game.Position().Status()
	if status == chess.NoMethod {
		if m.Game.Position().Turn() == playerColor {
			msg = "Il pubblico ha scelto!"
		} else {
			msg = "Il giocatore ha fatto la sua mossa, ora tocca al popolo!"
		}
	} else {
		if m.Game.Position().Turn() == playerColor {
			msg = "Il pubblico ha vinto! (per scacco o forfeit)"
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

func (m *Match) ASCII() string {
	str := m.Game.Position().Board().Draw()
	strings.ReplaceAll(str, "-", " - ")
	strings.ReplaceAll(str, "\n\r", "\n")
	return str
}

type SerializedMatch struct {
	Code     string          `json:"code"`
	Duration time.Duration   `json:"duration"`
	EndsAt   time.Time       `json:"ends_at"`
	Game     string          `json:"game"`
	Tweets   []request.Tweet `json:"tweets"`
}

func (m *Match) Serialized() SerializedMatch {
	return SerializedMatch{
		Code:     m.Code,
		Duration: m.Duration,
		EndsAt:   m.EndsAt,
		Game:     m.Game.FEN(),
		Tweets:   m.Tweets,
	}
}

func (m *Match) sendUpdate() {
	m.updates <- true
}

func (m *Match) Update() bool {
	return <-m.updates
}

func SetMatch(m *Match) {
	match = m
}

func GetMatch() *Match {
	return match
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func code(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (m *Match) setup() {
	m.updates = make(chan bool)
	m.quit = make(chan bool)
	m.delay()
	m.periodic()
}

func (m *Match) close() {
	m.quit <- false
	m.updates <- false

	if m.Game.Position().Turn() == playerColor {
		log.Println("Game closed. The mater lost")
	} else {
		log.Println("Game closed. The crowd lost")
	}
}

func NewMatch(duration time.Duration) *Match {
	m := Match{
		Code:     code(6),
		Duration: duration,
		EndsAt:   time.Now().Add(duration).UTC(),

		Game: chess.NewGame(),
	}

	m.setup()
	return &m
}
