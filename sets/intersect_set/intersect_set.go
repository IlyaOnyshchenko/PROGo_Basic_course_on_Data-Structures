// Intersect - пересечение множеств
func (s SetOnMap[T]) Intersect(otherSet SetOnMap[T]) {
	for item, _ := range s {
		_, ok := otherSet[item]
		if ok == false {
			s.Remove(item)
		}
	}
}