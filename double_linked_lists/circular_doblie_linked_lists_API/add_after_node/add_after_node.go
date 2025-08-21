// AddAfter - вставляет новый узел после переданного узла
func (list *DoubleCircularLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	newNode := &Node{
		Value:    item,
		Next:     node.Next,
		Previous: node,
	}
	node.Next.Previous = newNode
	node.Next = newNode
	list.count++
}