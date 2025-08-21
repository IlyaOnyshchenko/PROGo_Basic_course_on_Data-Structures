// Difference - разность множеств
func (s SetOnMap[T]) Difference(otherSet SetOnMap[T]) {
	for item, _ := range otherSet {
		delete(s, item)
	}
}