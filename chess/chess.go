package chess

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
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
}

func (m *Match) delay() {
	if m.ticking.Load() {
		m.timeout <- true
	}
	m.timeout = make(chan bool)

	go func() {
		m.ticking.Store(true)
		select {
		case <-time.After(match.Duration):
			m.ticking.Store(false)
			m.onTurnEnd()

		case <-match.timeout:
			m.ticking.Store(false)
			log.Println("Timeout cancled")
		}
	}()
}

func (m *Match) getMoves() (moves map[string]uint, err error) {
	tweets, err := request.Replies(m.TweetID, 100, "", "")
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		clone := match.Game.Clone()
		first := strings.Split(tweet.Text, " ")[0]
		if err = clone.MoveStr(first); err != nil {
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
		for move, val := range moves {
			if val > mostValued {
				mostRated = move
				mostValued = val
			}
		}
		m.Game.MoveStr(mostRated)
	}
}

func (m *Match) Move(move string) error {
	if match.Game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	if err := match.Game.MoveStr(move); err != nil {
		return err
	}

	m.PostGame()
	m.EndsAt = time.Now().UTC().Add(m.Duration)
	m.delay()
	return nil
}

func (m *Match) PostGame() {
	image, err := m.Image()
	if err != nil {
		log.Printf("WARN: Could not generate chessboard picture: %v", err)
		return
	}

	media, err := request.UploadMedia(image)
	if err != nil {
		log.Printf("WARN: Could not upload image to twitter: %v", err)
		return
	}
	tweet, err := request.Post(request.TweetRequest{
		Text: "bla bla bla",
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

func (m *Match) Image() (buf []byte, err error) {
	dest := bytes.NewBuffer(buf)
	if err = image.SVG(dest, m.Game.Position().Board()); err != nil {
		return
	}

	return dest.Bytes(), nil
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

func SetMatch(m *Match) {
	match = m
	if err := Store(); err != nil {
		log.Printf("Could not store chess match state: %v", err)
	}
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
	go func() {
		<-m.timeout
	}()

	m.delay()
	return m
}
