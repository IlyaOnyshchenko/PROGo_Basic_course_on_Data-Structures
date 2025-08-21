// RemoveFirstByValue - удаляет первый узел из списка, равный конкретному значению
func (dll *DoubleLinkedList) RemoveFirstByValue(value int) bool {
	current := dll.Head
	if current != nil {
		if current == dll.Head && current.Value == value {
			dll.RemoveFirst()
			return true
		}
		if current.Next == nil && current.Value == value {
			dll.RemoveLast()
			return true
		}
		current = current.Next
	}
	node := dll.Find(value)
	if node == nil {
		return false
	}
	prev := node.Previous
	prev.Next = node.Next
	if node.Next != nil {
		node.Next.Previous = prev
		dll.count--
		return true
	}
	return false
}