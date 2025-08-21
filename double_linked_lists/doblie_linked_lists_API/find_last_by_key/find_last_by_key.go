// FindLast - ищет последний узел по значению
func (dll *DoubleLinkedList) FindLast(key int) *Node {
	if dll.count == 0 {
		return nil
	}
	current := dll.Head
	var lastNode *Node
	for current != nil {
		if current.Value == key {
			lastNode = current
		}
		current = current.Next
	}
	return lastNode
}