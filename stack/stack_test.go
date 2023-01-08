package stack

import "testing"

func TestStackInt(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}

	if v, err := s.Top(); v != 1 || err != nil {
		t.Errorf("expected 1 but got: %v", v)
	}

	t.Log(s.Pop())

	if !s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}
}

func TestStackString(t *testing.T) {
	s := New[string]()
	if !s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}

	s.Push("abcd")
	if s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}

	if v, err := s.Top(); v != "abcd" || err != nil {
		t.Errorf("expected abcd but got: %v", v)
	}

	t.Log(s.Pop())

	if !s.IsEmpty() {
		t.Errorf("expected false but got: %v", s.IsEmpty())
	}
}
