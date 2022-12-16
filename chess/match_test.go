package chess

import (
	"fmt"
	"testing"
	"time"

	"github.com/notnil/chess"
	"github.com/stretchr/testify/assert"
)

func TestMatchCode(t *testing.T) {
	assert.Equal(t, "", code(0), "code(0) should return an empty string")
	for i := 0; i < 10; i++ {
		assert.Len(t, code(i), i, fmt.Sprintf("code(%d) should return a string of length %d", i, i))
	}
}

func TestInternalStart(t *testing.T) {
	m := Match{
		Duration: time.Minute,
		EndsAt:   time.Now().Add(time.Minute).UTC(),
		cancel:   make(chan bool),
		quit:     make(chan bool),
	}
	m.start()
	<-time.After(time.Second)
	assert.True(t, m.ticking.Load(), "Expected to find recurring services operational")
	m.end()
}

// TODO: fix
// func TestInternalEnd(t *testing.T) {
// 	m := Match{
// 		Duration: time.Minute,
// 		EndsAt:   time.Now().Add(time.Minute).UTC(),
// 		cancel:   make(chan bool),
// 		quit:     make(chan bool),
// 	}
// 	m.start()
// 	m.end()
// 	<-time.After(time.Second)
// 	assert.False(t, m.ticking.Load(), "Expected to find ticking stopped")
// }

func TestNewMatch(t *testing.T) {
	m := NewMatch(time.Minute)
	assert.NotNil(t, m, "Expected NewMatch to return a new match")
	assert.NotEmpty(t, m.Code, "Expected the new mach token to be non-empty")
	assert.Len(t, m.Code, 6, "Expected the new mach token to of length TOKEN_LEN")
	assert.Equal(t, m.Duration, time.Minute, "Expected the new mach to have the appropriate duration")
	assert.True(t, time.Now().Add(time.Minute).After(m.EndsAt), "Expected the new mach to have the appropriate ends at")
	assert.NotNil(t, m.Game, "Expected the new mach to have a not-nill gochess state")
	assert.Equal(t, m.Game.FEN(), chess.NewGame().FEN(), "Expected the new mach to have a not-nill gochess state")
	assert.False(t, m.Forfeited, "Expected the new mach not to be forfeited")

	m.end()
}

func TestSetMatch(t *testing.T) {
	SetMatch(nil)
	assert.Nil(t, match, "Expecteeeed the mtatch to be nil")
	SetMatch(NewMatch(time.Minute))
	assert.NotNil(t, match, "Expecteeeed the mtatch to be nil")
	GetMatch().end()
}

func TestGetMatch(t *testing.T) {
	SetMatch(nil)
	assert.Nil(t, GetMatch(), "Expecteeeed the mtatch to be nil")
	SetMatch(NewMatch(time.Minute))
	assert.NotNil(t, GetMatch(), "Expecteeeed the mtatch to be nil")
	GetMatch().end()
}

func TestSerializedMatch(t *testing.T) {
	m := NewMatch(time.Minute)
	sm := m.Serialized()
	assert.Equal(t, m.Code, sm.Code, "Expected match codes to match")
	assert.Equal(t, m.Duration, sm.Duration, "Expected match durations to match")
	assert.Equal(t, m.EndsAt, sm.EndsAt, "Expected match ends_at to match")
	assert.Equal(t, m.Game.FEN(), sm.Game, "Expected serialized game to be a the FEN representation")
	assert.Equal(t, m.Tweets, sm.Tweets, "Expected match tweets to be the same")
	assert.Equal(t, m.Forfeited, sm.Forfeited, "Expected forfeit status to match")
	m.end()
}
