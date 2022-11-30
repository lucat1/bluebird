package chess

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"log"
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
	playerColor = chess.White
	stateFile   = "chess.json"
)

var match *Match = nil

type Match struct {
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	EndsAt   time.Time     `json:"ends_at"`
	Game     *chess.Game   `json:"game"`
	TweetID  string        `json:"tweet_id"`

	timeout chan bool
	ticking atomic.Bool

	updates chan bool
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

		case <-time.After(m.Duration):
			m.ticking.Store(false)
			m.onTurnEnd()
		}
	}()
}

func (m *Match) getMoves() (moves map[string]uint, err error) {
	moves = map[string]uint{}
	tweets, err := request.Replies(m.TweetID, 100, "", "")
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		clone := m.Game.Clone()
		first := strings.Split(tweet.Text, " ")[1]
		if err := clone.MoveStr(first); err == nil {
			moves[first]++
		}
	}
	return
}

func (m *Match) onTurnEnd() {
	if m.Game.Position().Turn() == playerColor {
		// TODO: forfeit
	} else {
		moves, err := m.getMoves()
		if err != nil {
			log.Printf("Could not get tweets replies: %v", err)
			return
		}
		var (
			mostRated  string
			mostValued uint
		)
		if len(moves) != 0 {
			for move, val := range moves {
				if val > mostValued {
					mostRated = move
					mostValued = val
				}
			}
		} else {
			moves := m.Game.Moves()
			// TODO: if no random moves are available, game lost
			randMove := moves[rand.Intn(len(moves))]
			mostRated = randMove.String()
			log.Printf("No move given, playing random")
		}
		if err := m.Move(mostRated); err != nil {
			log.Printf("WARN: move %s failed:  %v", mostRated, err)
		}
	}
}

func (m *Match) Move(move string) error {
	if err := m.Game.MoveStr(move); err != nil {
		return err
	}

	log.Printf("Moving %s", move)
	m.updates <- true
	m.PostGame()
	m.EndsAt = time.Now().UTC().Add(m.Duration)
	m.delay()
	return Store()
}

func (m *Match) PlayerMove(move string) error {
	if m.Game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	return m.Move(move)
}

func (m *Match) PostGame() {
	var msg string
	if m.Game.Position().Turn() == playerColor {
		msg = "Il pubblico ha scelto!"
	} else {
		msg = "Il giocatore ha fatto la sua mossa, ora tocca al popolo!"
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
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	EndsAt   time.Time     `json:"ends_at"`
	Game     string        `json:"game"`
}

func (m Match) Serialized() SerializedMatch {
	return SerializedMatch{
		Code:     m.Code,
		Duration: m.Duration,
		EndsAt:   m.EndsAt,
		Game:     m.Game.FEN(),
	}
}

func (m Match) Update() {
	<-m.updates
}

func Store() (err error) {
	buf, err := json.Marshal(match)
	err = ioutil.WriteFile(stateFile, buf, 0666)
	return
}

func Resume() (err error) {
	var m Match
	buf, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return
	}
	if err = json.Unmarshal(buf, &m); err != nil {
		return err
	}
	match = &m
	return
}

func SetMatch(m *Match) error {
	match = m
	if err := Store(); err != nil {
		log.Printf("Could not store chess match state: %v", err)
	}
	return Store()
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

func NewMatch(duration time.Duration) Match {
	m := Match{
		Code:     code(6),
		Duration: duration,
		EndsAt:   time.Now().Add(duration).UTC(),

		Game: chess.NewGame(),
	}
	m.updates = make(chan bool)

	m.delay()
	return m
}
