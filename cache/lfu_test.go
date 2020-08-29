package cache

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
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
	lfu.Add("key", "something")
	lfu.Add("key2", 32) //int type
	v, err := lfu.Get("key")
	assert.Nil(t, err)
	fmt.Println("VALUE:", v)
	spew.Dump(lfu)
}
