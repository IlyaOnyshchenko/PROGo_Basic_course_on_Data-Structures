// Remove - удаляет узел по ключу
func (l *SinglyLinkedList) Remove(key int) bool {
	if l.count == 0 {
		return false
	}
	next := l.Find(key)
	if next == nil {
		return false
	}
	current := l.Head
	var previous *Node
	for current != nil {
		if current.Next == next {
			previous = current
			previous.Next = next.Next
		}
		if current == next {
			l.Head = current.Next
			l.count--
			return true
		}
		current = current.Next
	}
	l.count--
	if previous.Next == next.Next {
		return true
	} else {
		return false
	}
}