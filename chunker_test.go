package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestShouldChunkText(t *testing.T) {
	chunks := ToChunks("1234567890", 4)
	fmt.Println(FromChunks(chunks))
	assert.True(t, "1234567890" == FromChunks(chunks))
}
