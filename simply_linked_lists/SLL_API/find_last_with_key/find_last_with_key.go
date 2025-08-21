// FindLast - ищет последний узел равный переданному значению
func (l *SinglyLinkedList) FindLast(key int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}
	current := l.Head
	var current1 *Node
	for current != nil {
		if current.Value == key {
			current1 = current
		}
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current1
}