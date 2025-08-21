func (s SetOnMap[T]) Union(otherSet SetOnMap[T]) {
	for k, v := range otherSet {
		s[k] = v
	}
}