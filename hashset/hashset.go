package hashset


type HashSet[T comparable] struct {
  elements map[T]bool
}

func New[T comparable]() HashSet[T] {
  return HashSet[T]{make(map[T]bool)}
}

func (s HashSet[T]) Add(element T) bool {
  if _, ok := s.elements[element]; ok {
    return false
  }
  s.elements[element] = true
  return true
}

func (s HashSet[T]) Clear() {
  s.elements = make(map[T]bool)
}

func (s HashSet[T]) Contains(element T) bool {
  return s.elements[element]
}

func (s HashSet[T]) Empty() bool {
  return s.Size() == 0 
}

func (s HashSet[T]) Elements() []T {
  elements := make([]T, 0, len(s.elements))
  for element, _ := range s.elements {
    elements = append(elements, element)
  }
  return elements 
}

func (s HashSet[T]) Remove(element T) bool {
  if _, ok := s.elements[element]; !ok {
    return false
  }
  delete(s.elements, element)
  return true
}

func (s HashSet[T]) Size() int {
  return len(s.elements)
}
