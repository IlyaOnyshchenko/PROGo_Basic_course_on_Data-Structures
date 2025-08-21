// AddAfter - вставляет узел в список после переданного
func (dll *DoubleLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	if node.Next == nil {
		dll.PushBack(item)
		return
	}
	newNode := &Node{
		Value:    item,
		Next:     node.Next,
		Previous: node,
	}
	node.Next.Previous = newNode
	node.Next = newNode
	dll.count++
}