package cache

import (
	"fmt"
	LIST "vrecan/data-structures/list"
)

//ErrNotExist doesn't exist
var ErrNotExist = fmt.Errorf("does not exist")

//LFU least frequently used cache.
type LFU struct {
	cache   map[string]interface{}
	list    LIST.DoublyLinkedList
	maxSize int
}

//Entry to the LFU
type Entry struct {
	Value     interface{}
	ValueNode *LIST.Node
}

//NewLFU create a new least frequentlly used cache
func NewLFU(maxSize int) *LFU {
	return &LFU{
		cache:   make(map[string]interface{}, 0),
		list:    LIST.DoublyLinkedList{},
		maxSize: maxSize,
	}
}

//Add an item to the LFU
func (l *LFU) Add(key string, value interface{}) {
	_, ok := l.cache[key]
	if !ok {
		l.addNewEntry(key, value)
	}
}

func (l *LFU) addNewEntry(key string, value interface{}) {
	fmt.Println("add entry!")
	if l.list.Head != nil && l.list.Head.Value == 1 {
		l.cache[key] = Entry{Value: value, ValueNode: l.list.Head}
		return
	}
	//first node is nil or not == 1 so we need to create a new node
	l.list.Push(&LIST.Node{Value: 1})
	l.cache[key] = Entry{Value: value, ValueNode: l.list.Head}
	return

}

//Get value from cache by key
func (l *LFU) Get(key string) (value interface{}, err error) {
	v, ok := l.cache["key"]
	if !ok {
		return nil, ErrNotExist
	}
	entry := v.(Entry)
	//end of the list
	if entry.ValueNode.Next == nil {
		//if no next add it with the new value
		newNode := &LIST.Node{Value: entry.ValueNode.Value.(int) + 1,
			Pre: entry.ValueNode}
		l.list.Append(newNode)
		entry.ValueNode = newNode
	}
	// Test for next value and if it exists swap to that one

	// create new entry if not true changing the pointers for next and previous

	//keep track of how many nodes point to a frequency so we can delete them when nothing
	//points to them anymore

	return v, nil
}
