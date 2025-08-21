// PushBackRange - добавляет список узлов в конец
func (l *SinglyLinkedList) PushBackRange(items []int) {
	tail := l.FindTail()
	for i := 0; i < len(items); i++ {
		newNode := &Node{Value: items[i]}
		if l.count == 0 {
			l.Head = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		l.count++
	}
}