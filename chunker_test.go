package chunker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldChunkText(t *testing.T) {
	str := "sdpoifgj npowi4uvtnp349t8cnw;sfgimxao;ijncw;omgw;ot3i5vngw;p3o4jmvsehrngvs"

	chunks := ToChunks(str, 4)

	assert.True(t, str == FromChunks(chunks))
}
