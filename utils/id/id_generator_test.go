package id

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	assert := assert.New(t)
	id1 := New()
	time.Sleep(1 * time.Second)
	id2 := New()
	assert.NotEqual(id1, id2)
	assert.Equal(len(id1), 20)
	assert.Equal(len(id2), 20)
}
