package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomSentiment (t *testing.T) {

  A := randomSentiment()
  
  assert.Equal(t,A[0]+A[1]+A[2]+A[3],float32(1),"Random Sentiment Values do not add up to 1")

}
