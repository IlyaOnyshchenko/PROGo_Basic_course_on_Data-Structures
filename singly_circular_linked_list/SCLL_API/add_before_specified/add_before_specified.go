// AddBefore - добавляет узел до переданного
func (list *SinglyCircularLinkedList) AddBefore(node *Node, item int) {
	newNode := &Node{Value: item}
	var previous *Node
	current := list.Tail
	for {
		if current.Next == node {
			previous = current
		}
		current = current.Next
		if current == list.Tail {
			break
		}
	}
	previous.Next = newNode
	newNode.Next = node
	list.count++
}