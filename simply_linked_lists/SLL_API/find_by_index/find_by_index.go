// FindByIndex - ищет узел списка по индексу
func (l *SinglyLinkedList) FindByIndex(index int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}
	current := l.Head
	cnt := 0
	for current != nil {
		if cnt == index {
			return current
		}
		current = current.Next
		cnt++
	}
	return nil
}