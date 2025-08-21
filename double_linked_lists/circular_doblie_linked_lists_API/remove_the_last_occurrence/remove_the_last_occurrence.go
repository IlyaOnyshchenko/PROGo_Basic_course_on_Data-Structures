// RemoveLastByValue - удаление последнего найденного узла по значению
func (list *DoubleCircularLinkedList) RemoveLastByValue(item int) bool {
	node := list.FindLast(item)
	if node == nil {
		return false
	}
	if node.Next == node {
		list.Head = nil
		list.count--
		return true
	} else {
		node.Next.Previous = node.Previous
		node.Previous.Next = node.Next
		if list.Head == node {
			list.Head = node.Next
		}
		list.count--
		return true
	}
	return false
}
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