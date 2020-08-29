package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLFU(t *testing.T) {
	lfu := NewLFU(23)
	assert.Equal(t, lfu.maxSize, 23)
}

func TestBasicLFU(t *testing.T) {
	lfu := NewLFU(23)
	assert.Equal(t, lfu.maxSize, 23)

	lfu.Add("key", "something")
	lfu.Add("key2", 32) //int type
}
