package chess

import (
	"os"
	"testing"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"git.hjkl.gq/team14/team14/test"
	"github.com/notnil/chess"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	postImageServer := test.CreateMultiServer(map[string][]byte{
		"/": test.ReadFile("../mock/upload.json"),
	})
	defer postImageServer.Close()
	postPostServer := test.CreateMultiServer(map[string][]byte{
		"/tweets": test.ReadFile("../mock/post.json"),
	})
	defer postPostServer.Close()
	client, err := request.NewV1Client(postImageServer.URL, postPostServer.URL, "", "", "", "")
	if err != nil {
		panic(err)
	}
	request.SetV1Client(client)
	code := m.Run()
	os.Exit(code)
}

func TestForfeit(t *testing.T) {
	m := NewMatch(time.Minute)
	assert.False(t, m.Forfeited, "Expected new match not to be forfeited")
	m.Forfeit()
	assert.True(t, m.Forfeited, "Expected the forfeited property to have been set")
	assert.Equal(t, chess.BlackWon, m.Game.Outcome(), "Expected the black to have won")
	assert.Equal(t, chess.Resignation, m.Game.Method(), "Expected the game to have been resigned")
	assert.Equal(t, chess.Resignation, m.Game.Method(), "Expected the game to have been resigned")
}
