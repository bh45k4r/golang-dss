package node

type Node[T any] struct {
  value T
  prev *Node[T]
  next *Node[T]
}

func New[T any](value T, prev, next *Node[T]) *Node[T] {
  return &Node[T]{value, prev, next}
}

func (n *Node[T]) Value() T {
  return n.value
}

func (n *Node[T]) GetPrev() *Node[T] {
  return n.prev
}

func (n *Node[T]) GetNext() *Node[T] {
  return n.next
}

func (n *Node[T]) SetPrev(prev *Node[T]) {
  n.prev = prev
}

func (n *Node[T]) SetNext(next *Node[T]) {
  n.next = next
}
