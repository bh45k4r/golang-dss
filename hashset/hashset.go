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

func (s HashSet[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func(ch chan<- T) {
		defer close(ch)
		for k, _ := range s.elements {
			ch <- k
		}
	}(ch)
	return ch
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
