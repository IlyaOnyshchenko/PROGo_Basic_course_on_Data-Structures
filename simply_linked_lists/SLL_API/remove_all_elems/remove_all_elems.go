// RemoveAll - удаляет все узлы по ключу
func (l *SinglyLinkedList) RemoveAll(key int) int {
	if l.count == 0 {
		return 0
	}
	cnt := 0
	current := l.Head
	for current != nil {
		if current.Value == key {
			l.Remove(key)
			cnt++
		}
		current = current.Next
	}
	return cnt
}