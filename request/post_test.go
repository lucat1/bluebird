package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	SetV1Client(postMediaClient)
	media := []byte{0, 0, 0, 0, 0}
	mres, err := PostMedia(media)

	assert.Nil(t, err, "Expected no error in function call PostMedia")
	assert.Equal(t, mres.MediaID, "1604127079230017536", "Expected equal value in res")

	IDs := TweetRequestMedia{[]string{mres.MediaID}}
	req := TweetRequest{"test", IDs}

	cres, err := PostCustom(req)
	assert.Nil(t, err, "Expected no error in function call PostCustom")
	assert.Equal(t, cres.ID, "1604169175437639686", "Expected IDs to match")
}
