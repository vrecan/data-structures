package list

//DoublyLinkedList is a list that points to the head and tail
type DoublyLinkedList struct {
	Head   *Node // The first node
	Tail   *Node // The last node
	Length int   // The number of nodes
}

//Node is a generic data structure for anything stored in the list
type Node struct {
	Value interface{}
	Pre   *Node
	Next  *Node
}

//Append to the list
func (dl *DoublyLinkedList) Append(n *Node) {
	if n == nil {
		return
	}
	// Check if empty
	if dl.Length == 0 {
		dl.Head = n
		dl.Tail = n
		dl.Length++
		return
	}

	lastNode := dl.Tail

	// Then, we update the last node (tail) in our Doubly linked list
	// with the new Node, connecting it to the previous last Node
	// and also connecting our previous last Node to the new one.
	// Finally, we only need to increment the length.
	dl.Tail = n
	dl.Tail.Pre = lastNode
	lastNode.Next = n
	dl.Length++
}
