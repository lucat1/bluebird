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

var (
	noMovesClient *request.RequestClient
	movesClient   *request.RequestClient
)

func TestMain(m *testing.M) {
	var err error
	movesServer := test.CreateMultiServer(map[string][]byte{
		"/tweets/search/recent": test.ReadFile("../mock/by_convid_moves.json"),
		"/tweets":               test.ReadFile("../mock/by_tweetid.json"),
	})
	defer movesServer.Close()
	noMovesServer := test.CreateMultiServer(map[string][]byte{
		"/tweets/search/recent": test.ReadFile("../mock/by_convid.json"),
		"/tweets":               test.ReadFile("../mock/by_tweetid.json"),
	})
	defer movesServer.Close()
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
	movesClient, err = request.NewClient(movesServer.URL, "")
	if err != nil {
		panic(err)
	}
	noMovesClient, err = request.NewClient(noMovesServer.URL, "")
	if err != nil {
		panic(err)
	}
	request.SetClient(movesClient)
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

func TestOnTurnEndNoWhiteMove(t *testing.T) {
	m := NewMatch(time.Second)
	<-time.After(time.Second*time.Duration(1) + time.Millisecond*time.Duration(100))
	assert.Equal(t, m.Game.Position().Turn(), chess.Black, "Expected the player turn to be black after a random move")
	m.end()
}

func TestOnTurnEndNoBlackMove(t *testing.T) {
	request.SetClient(noMovesClient)
	m := NewMatch(time.Second)
	m.PlayerMove("d3")
	<-time.After(time.Second*time.Duration(1) + time.Millisecond*time.Duration(100))
	assert.Equal(t, m.Game.Position().Turn(), chess.White, "Expected the player turn to be white after a random move (from the crowd)")
	m.end()
	request.SetClient(movesClient)
}

func TestOnTurnEndBlackMoved(t *testing.T) {
	m := NewMatch(time.Second)
	m.PlayerMove("d3")
	<-time.After(time.Second*time.Duration(1) + time.Millisecond*time.Duration(100))
	assert.Equal(t, m.Game.Position().Turn(), chess.White, "Expected the player turn to be white after a random move (from the crowd)")

	exp := chess.NewGame()
	exp.MoveStr("d3")
	exp.MoveStr("f6")
	assert.Equal(t, exp.FEN(), m.Game.FEN(), "Expected the moved to have been played")
	m.end()
}
