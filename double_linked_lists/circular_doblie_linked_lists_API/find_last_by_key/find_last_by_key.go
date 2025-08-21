// FindLast - находит последний узел по значению
func (l *DoubleCircularLinkedList) FindLast(key int) *Node {
	if l.count == 0 {
		return nil
	}
	if l.Find == nil {
		return nil
	}
	current := l.Head
	var lastNode *Node
	for {
		if current.Value == key {
			lastNode = current
		}
		current = current.Next
		if current == l.Head {
			break
		}
	}
	return lastNode
}