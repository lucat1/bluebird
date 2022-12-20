package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWord(t *testing.T) {
	myurl := "https://pbs.twimg.com/media/FkXM2IyXgAAx-9t?format=jpg&name=small"
	SetOCRClient(ocrTestClient)
	word, err := GetWord(myurl)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "PALLONE", word, "Expected word was PALLONE")

	SetOCRClient(ocrErrorClient)
	word, err = GetWord(myurl)
	assert.NotNil(t, err, "Expected error")

	SetOCRClient(ocrJSONErrorClient)
	word, err = GetWord(myurl)
	assert.NotNil(t, err, "Expected error")
}
