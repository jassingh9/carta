package carta

// The animalSet type is a type alias of `map[string]struct{}`
type Set[T comparable] map[T]struct{}

// Adds an animal to the set
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Removes an animal from the set
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Returns a boolean value describing if the animal exists in the set
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}
