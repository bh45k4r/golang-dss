package stack

import (
	"errors"
	"github.com/bh45k4r/golang-dss/node"
)

type Stack[T comparable] struct {
	head *node.Node[T]
}

func New[T comparable]() *Stack[T] {
  return &Stack[T]{nil}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Pop() (T, error) {
  var value T
	if s.IsEmpty() {
		return value, errors.New("stack is empty")
	}
	value, s.head = s.head.Value(), s.head.GetNext()
	return value, nil
}

func (s *Stack[T]) Push(value T) {
	node := node.New(value, nil, nil) 
	if s.IsEmpty() {
		s.head = node
		return
	}
	node.SetNext(s.head)
  s.head = node
}

func (s *Stack[T]) Top() (T, error) {
  var value T
	if s.IsEmpty() {
		return value, errors.New("stack is empty")
	}
	return s.head.Value(), nil
}
