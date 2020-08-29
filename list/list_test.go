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

func TestAppendAfterHead(t *testing.T) {

	list := DoublyLinkedList{}

	total := 10

	for i := 1; i <= total; i++ {
		n := &Node{Value: i}
		list.Push(n)
	}
	assert.Equal(t, list.Length, 10)
	newNode := &Node{Value: "new"}
	expNext := list.Head.Next
	err := list.AppendAfter(newNode, list.Head)
	assert.Nil(t, err)
	assert.Equal(t, list.Head.Next, newNode)
	assert.Equal(t, newNode.Next, expNext)
}

func TestAppendAfterTail(t *testing.T) {

	list := DoublyLinkedList{}

	total := 10

	for i := 1; i <= total; i++ {
		n := &Node{Value: i}
		list.Push(n)
	}
	assert.Equal(t, list.Length, 10)
	newNode := &Node{Value: "new"}
	expNext := list.Tail.Next
	err := list.AppendAfter(newNode, list.Tail)
	assert.Nil(t, err)
	assert.Equal(t, list.Tail, newNode)
	assert.Equal(t, newNode.Next, expNext)
	spew.Dump(list)
}
