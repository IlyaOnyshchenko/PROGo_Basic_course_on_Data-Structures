// Remove - удаляет первое вхождение указанного элемента из списка
func (l *SinglyCircularLinkedList) Remove(key int) bool {
	if l.count == 0 {
		return false
	}
	next := l.Find(key)
	if l.Find(key) == nil {
		return false
	}
	current := next
	for current.Next != next {
		current = current.Next
	}
	l.removeAfterNodeInternal(current)
	return true
}