package list

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {

	list := DoublyLinkedList{}

	for i := 1; i <= 10; i++ {
		n := &Node{Value: i}
		list.Append(n)
	}
	assert.Equal(t, list.Length, 10)
	node := list.Head
	expValue := 1
	expPrev := 1
	for node != nil {
		if expValue > 1 {
			assert.Equal(t, node.Pre.Value, expPrev)
			expPrev++
		}
		assert.Equal(t, node.Value, expValue)

		fmt.Println(node.Value)
		node = node.Next
		expValue++
	}
}

func TestPush(t *testing.T) {

	list := DoublyLinkedList{}

	total := 10

	for i := 1; i <= total; i++ {
		n := &Node{Value: i}
		list.Push(n)
	}
	assert.Equal(t, list.Length, 10)
	node := list.Head
	expValue := total
	expPrev := total
	spew.Dump(list)
	for node != nil {

		if list.Length != expValue {
			assert.Equal(t, node.Pre.Value, expPrev)
			expPrev--
		}

		assert.Equal(t, node.Value, expValue)

		fmt.Println(node.Value)
		node = node.Next
		expValue--
	}

}
