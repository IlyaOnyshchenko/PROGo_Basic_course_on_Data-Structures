// RemoveLastKey - удаляет последний узел по ключу
func (l *SinglyCircularLinkedList) RemoveLastKey(key int) bool {
	if l.count == 0 {
		return false
	}
	next := l.FindLast(key)
	if l.FindLast(key) == nil {
		return false
	}
	current := next
	for current.Next != next {
		current = current.Next
	}
	l.removeAfterNodeInternal(current)
	return true
}