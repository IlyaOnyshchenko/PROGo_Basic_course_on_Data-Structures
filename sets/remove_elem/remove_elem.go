func (s SetOnMap[T]) Remove(item T) bool {
	for i, _ := range s {
		if i == item {
			delete(s, i)
			return true
		}
	}
	return false
}