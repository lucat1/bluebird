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
	Code     string
	Duration time.Duration
	EndsAt   time.Time
	game     *chess.Game
	tweetID  string
	Tweets   []request.Tweet

	timeout chan bool
	ticking atomic.Bool

	fetching      atomic.Bool
	fetched       chan bool
	awaitingFetch atomic.Int32
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
	tweets, err := request.Replies(m.tweetID, 100, "", "")
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		clone := match.game.Clone()
		first := strings.Split(tweet.Text, " ")[0]
		if err = clone.MoveStr(first); err != nil {
			moves[first]++
		}
	}
	return
}

func (m *Match) onTurnEnd() {
	if m.game.Position().Turn() == playerColor {
		// TODO: forfeit
	} else {
		m.fetching.Store(true)
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
		m.game.MoveStr(mostRated)
		m.fetched <- true
		m.fetching.Store(false)
	}
}

func (m *Match) Move(move string) error {
	if match.game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	if err := match.game.MoveStr(move); err != nil {
		return err
	}

	m.EndsAt = time.Now().UTC().Add(m.Duration)
	m.delay()
	return nil
}

func (m *Match) Image() (buf []byte, err error) {
	dest := bytes.NewBuffer(buf)
	if err = image.SVG(dest, m.game.Position().Board()); err != nil {
		return
	}

	return dest.Bytes(), nil
}

type SerializedMatch struct {
	Code     string          `json:"code"`
	Duration time.Duration   `json:"duration"`
	EndsAt   time.Time       `json:"ends_at"`
	Game     string          `json:"game"`
	Tweets   []request.Tweet `json:"tweets"`
}

func (m Match) Serialized() SerializedMatch {
	return SerializedMatch{
		Code:     m.Code,
		Duration: m.Duration,
		EndsAt:   m.EndsAt,
		Game:     m.game.FEN(),
		Tweets:   m.Tweets,
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
	if match != nil {
		if match.fetching.Load() {
			match.awaitingFetch.Add(1)
			<-match.fetched
		}
	}
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

		game: chess.NewGame(),
	}
	go func() {
		<-m.timeout
	}()

	m.delay()
	return m
}
