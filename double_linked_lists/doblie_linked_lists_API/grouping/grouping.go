// Group - группирует узлы списка так, что сначала идут узлы не четных позиций, затем четных
func (dll *DoubleLinkedList) Group() {
	oddList := &DoubleLinkedList{}
	evenList := &DoubleLinkedList{}
	current := dll.Head
	index := 1
	for current != nil {
		next := current.Next
		if index%2 == 0 {
			evenList.PushBack(current.Value)
		} else {
			oddList.PushBack(current.Value)
		}
		current = next
		index++
	}
	tail := oddList.findTail()
	if tail != nil {
		tail.Next = evenList.Head
		if evenList.Head != nil {
			evenList.Head.Previous = tail
		}
	}
	dll.Head = oddList.Head
}