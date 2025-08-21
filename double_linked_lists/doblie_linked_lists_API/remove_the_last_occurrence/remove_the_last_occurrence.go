// RemoveLastByValue - удаляет последний узел из списка, равный конкретному значению
func (dll *DoubleLinkedList) RemoveLastByValue(value int) bool {
	current := dll.Head
	var lastNode *Node
	for current != nil {
		if current.Value == value {
			lastNode = current
		}
		current = current.Next
	}
	if lastNode == nil {
		return false
	}
	if lastNode.Next == nil {
		dll.RemoveLast()
		return true
	}
	if dll.Head == lastNode {
		dll.RemoveFirst()
		return true
	}
	prev := lastNode.Previous
	prev.Next = lastNode.Next
	if lastNode.Next != nil {
		lastNode.Next.Previous = prev
		dll.count--
		return true
	}
	return false
}