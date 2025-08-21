func (l *SinglyLinkedList) AddBefore(node *Node, item int) {
	if node == nil || l.Head == nil {
		return
	}
	newNode := &Node{Value: item}
	newNode.Next = node
	if node == l.Head {
		l.Head = newNode
		l.count++
		return
	} else if newNode == l.Head {
		newNode.Next = l.Head
		l.Head = newNode
		l.count++
		return
	}
	current := l.Head
	var previous *Node
	for current != nil {
		if current.Next == node {
			previous = current
			break
		}
		current = current.Next
	}
	if current == nil {
		return
	}
	previous.Next = newNode
	l.count++
}