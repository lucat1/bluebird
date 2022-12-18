package chess

import (
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"github.com/notnil/chess"
)

var match *Match = nil

const (
	playerColor       = chess.White
	ReplyPollInterval = 15 // seconds
	TokenLen          = 6
)

type Match struct {
	Code      string
	Duration  time.Duration
	EndsAt    time.Time
	Game      *chess.Game
	TweetID   string
	Tweets    []request.Tweet
	Moves     map[string]uint
	Forfeited bool

	// player move cancel
	cancel  chan bool
	ticking atomic.Bool

	// updates chanel
	updates chan bool

	// periodic
	quit chan bool
}

func (m *Match) start() {
	m.delay(m.onTurnEnd)
	m.periodic()
}

func (m *Match) end() {
	select {
	case m.cancel <- false:
	default:
	}
	select {
	case m.quit <- false:
	default:
	}
	select {
	case m.cancel <- false:
	default:
	}

	log.Println("Chess game ended")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func code(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewMatch(duration time.Duration) *Match {
	m := Match{
		Code:     code(TokenLen),
		Duration: duration,
		EndsAt:   time.Now().Add(duration).UTC(),

		Game:      chess.NewGame(),
		Forfeited: false,

		cancel:  make(chan bool),
		updates: make(chan bool),
		quit:    make(chan bool),
	}

	m.start()
	return &m
}

func SetMatch(m *Match) {
	match = m
}

func GetMatch() *Match {
	return match
}

type SerializedMatch struct {
	Code      string          `json:"code"`
	Duration  time.Duration   `json:"duration"`
	EndsAt    time.Time       `json:"ends_at"`
	Game      string          `json:"game"`
	Tweets    []request.Tweet `json:"tweets"`
	Moves     map[string]uint `json:"moves"`
	Forfeited bool            `json:"forfeited"`
}

func (m *Match) Serialized() SerializedMatch {
	return SerializedMatch{
		Code:      m.Code,
		Duration:  m.Duration,
		EndsAt:    m.EndsAt,
		Game:      m.Game.FEN(),
		Tweets:    m.Tweets,
		Moves:     m.Moves,
		Forfeited: m.Forfeited,
	}
}
