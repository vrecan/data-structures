package cache

import (
	LIST "vrecan/data-structures/list"
)

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
	if _, ok := l.cache[key]; ok {
		if !ok {
			l.addNewEntry(key, value)
		}
	}
}

func (l *LFU) addNewEntry(key string, value interface{}) {
	if l.list.Head != nil && l.list.Head.Value == 1 {
		l.cache[key] = Entry{Value: value, ValueNode: l.list.Head}
		return
	}
	//first node is nil or not == 1 so we need to create a new node
	l.list.Push(&LIST.Node{Value: 1})
	l.cache[key] = Entry{Value: value, ValueNode: l.list.Head}
	return

}
