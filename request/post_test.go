package request
import (
	"github.com/stretchr/testify/assert"
	"testing"
	
)

func TestPostMedia(t *testing.T){
	SetV1Client(postMediaClient)
	media := []byte{0,0,0,0,0} 
	res, err := PostMedia(media)

	assert.Nil(t, err, "Expected no error in function call PostMedia")
	assert.Equal(t, res.MediaID, "1604127079230017536", "Expected equal value in res")
}