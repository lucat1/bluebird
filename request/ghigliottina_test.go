package request

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGhigliottina(t *testing.T) {
	SetClient(ghigliottinaClient)
	startTime, _ := time.Parse(time.RFC3339, "2022-12-12T00:00:00.000Z")
	endTime, _ := time.Parse(time.RFC3339, "2022-12-14T00:00:00.000Z")
	res, err := Ghigliottina(&startTime, &endTime)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "CANTANTE", res.Word, "Expected word to guess was CANTANTE")
	assert.Equal(t, "H90CJNow", res.Podium.First.Username, "The first place should be H90CJNow")
	assert.Equal(t, "metantacauta72", res.Podium.Second.Username, "The second place should be metantacauta72")
	assert.Equal(t, "equuscavallo", res.Podium.Third.Username, "The third place should be equuscavallo")

	SetClient(ghigliottinaTweetsErrorClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.NotNil(t, err, "Expected an error")

	SetClient(ghigliottinaNoTweetsClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.Equal(t, errors.New("No tweets were found"), err, "Expected no tweets found")

	SetClient(ghigliottinaNoRepliesResponseClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.NotNil(t, err, "Expected an error")

	SetClient(ghigliottinaNoRepliesClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.Equal(t, errors.New("No tweet replies were found"), err, "Expected no replies found")

	SetClient(ghigliottinaWinnersErrorClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.Equal(t, errors.New("Winners are malformed"), err, "Expected an error while parsing winners")

	SetClient(ghigliottinaTimeErrorClient)
	res, err = Ghigliottina(&startTime, &endTime)
	assert.Equal(t, errors.New("Times are malformed"), err, "Expected an error while parsing times")
}
