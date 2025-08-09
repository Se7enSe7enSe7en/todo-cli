package set

/*
A Set is a collection of unique elements.
  - the set implementation is just a map with keys of string and values of empty struct
  - empty struct (struct{}) takes up no memory
*/
type Set struct {
	elements map[string]struct{}
}

// Creates a new set
func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

// Inserts an element into the set
func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

// Deletes an element from the set
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Checks if an element exist in the set
func (s *Set) Contains(value string) bool {
	_, isExist := s.elements[value]
	return isExist
}

// Gets the length of the set
func (s *Set) Size() int {
	return len(s.elements)
}

// Gets all the elements of the set and returns them in a slice
func (s *Set) List() []string {
	list := make([]string, len(s.elements))

	for key := range s.elements {
		list = append(list, key)
	}

	return list
}
