// IsSubset - проверяет, является ли otherSet подмножеством текущего множества
func (s SetOnMap[T]) IsSubset(otherSet SetOnMap[T]) bool {
	for item, _ := range otherSet {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}