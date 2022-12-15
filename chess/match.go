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
)

type Match struct {
	Code      string
	Duration  time.Duration
	EndsAt    time.Time
	Game      *chess.Game
	TweetID   string
	Tweets    []request.Tweet
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
		Code:     code(6),
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
