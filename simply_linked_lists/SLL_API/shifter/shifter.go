// Reverse - переворачивает все элементы в обратном порядке
func (l *SinglyLinkedList) Reverse() {
	if l.count == 0 {
		return
	}
	current := l.Head
	var previous *Node
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	l.Head = previous
}