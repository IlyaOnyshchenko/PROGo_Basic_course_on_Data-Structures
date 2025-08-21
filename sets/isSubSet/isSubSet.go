func (s *SetOnSlice[T]) IsSubset(otherSet *SetOnSlice[T]) bool {
	for _, item := range s.items {
		if otherSet.contains(item) == false {
			return false
		}
	}
	return true
}