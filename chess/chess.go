package chess

import (
	"encoding/json"
	"errors"
	"git.hjkl.gq/team14/team14/request"
	"github.com/notnil/chess"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
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
	timeout  chan bool
	tweetID  string
}

func delay(fn func()) {
	if match.timeout != nil {
		match.timeout <- true
	}
	match.timeout = make(chan bool)

	select {
	case <-time.After(match.Duration):
		fn()

	case <-match.timeout:
		log.Println("Timeout cancled")
	}
}

func (m *Match) checkMoves() (moves map[string]uint, err error) {
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

func (m *Match) Move(move string) error {
	if match.game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	if err := match.game.MoveStr(move); err != nil {
		return err
	}

	m.EndsAt = time.Now().UTC().Add(m.Duration)
	m.timeout <- true
	return nil
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
		Game:     m.game.FEN(),
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
	return Match{
		Code:     code(6),
		Duration: duration,
		EndsAt:   time.Now().Add(duration).UTC(),

		game: chess.NewGame(),
	}
}
