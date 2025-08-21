// FindLast - ищет последний узел в списке
func (list *SinglyCircularLinkedList) FindLast(item int) *Node {
	if list.count == 0 {
		return nil
	}
	current := list.Tail
	var reqNode *Node
	for {
		if current.Value == item {
			reqNode = current
		}
		if current.Value == item && current == list.Tail {
			return current
		}
		current = current.Next
		if current == list.Tail {
			break
		}
	}
	return reqNode
}