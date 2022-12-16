package chess

import (
	"testing"
	"time"

	"github.com/notnil/chess"
	"github.com/stretchr/testify/assert"
)

func TestForfeit(t *testing.T) {
	m := NewMatch(time.Minute)
	assert.False(t, m.Forfeited, "Expected new match not to be forfeited")
	m.Forfeit()
	assert.True(t, m.Forfeited, "Expected the forfeited property to have been set")
	assert.Equal(t, chess.BlackWon, m.Game.Outcome(), "Expected the black to have won")
	assert.Equal(t, chess.Resignation, m.Game.Method(), "Expected the game to have been resigned")
	assert.Equal(t, chess.Resignation, m.Game.Method(), "Expected the game to have been resigned")
}
