package chess

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Player string

const (
	PlayerMaster Player = "master"
	PlayerOthers        = "others"
)

type Match struct {
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	EndsAt   time.Time     `json:"ends_at"`

	// TODO: replace with the proper go chess handling library
	State string `json:"state"`
	Turn  Player `json:"turn"`
}

const stateFile = "chess.json"

var match *Match = nil

func Store() (err error) {
	buf, err := json.Marshal(match)
	err = ioutil.WriteFile(stateFile, buf, 0666)
	return
}

func Resume() (err error) {
	buf, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return
	}
	return json.Unmarshal(buf, match)
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
		EndsAt:   time.Now().UTC().Add(duration),

		Turn:  PlayerMaster,
		State: "moving",
	}
}
