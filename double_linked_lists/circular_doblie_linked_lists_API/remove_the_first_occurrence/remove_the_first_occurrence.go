// Remove - удаление первого найденного узла по значению
func (list *DoubleCircularLinkedList) Remove(item int) bool {
	node := list.Find(item)
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