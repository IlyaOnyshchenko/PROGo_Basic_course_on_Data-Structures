// RemoveLastKey - удаляет последний узел по ключу
func (l *SinglyLinkedList) RemoveLastKey(key int) bool {
	current := l.Head
	current1 := l.FindLast(key)
	if current1 == nil {
		return false
	}
	for current != nil {
		if current.Next == current1 {
			current.Next = current.Next.Next
			l.count--
			return true
		}
		if current == nil {
			return false
		}
		current = current.Next
	}
	return false
}