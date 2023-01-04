package node


import "testing"


func TestNodeInt(t *testing.T) {
  n1 := New(11, nil, nil)
  n2 := New(13, n1, nil)
  n1.SetNext(n2)

  if n1.Value() != 11 {
    t.Errorf("expected: 11, but got: %d", n1.Value())
  }
  if n2.Value() != 13 {
    t.Errorf("expected: 13, but got: %d", n2.Value())
  }

  if n1.GetNext() != n2 {
    t.Errorf("expected: %v, but got: %v", n2, n1.GetNext())
  }
  if n2.GetPrev() != n1 {
    t.Errorf("expected: %v, but got: %v", n1, n2.GetPrev())
  }
}


func TestNodeString(t *testing.T) {
  n1 := New("aa", nil, nil)
  n2 := New("bb", n1, nil)
  n1.SetNext(n2)

  if n1.Value() != "aa" {
    t.Errorf("expected: aa, but got: %s", n1.Value())
  }
  if n2.Value() != "bb" {
    t.Errorf("expected: bb, but got: %s", n2.Value())
  }

  if n1.GetNext() != n2 {
    t.Errorf("expected: %v, but got: %v", n2, n1.GetNext())
  }
  if n2.GetPrev() != n1 {
    t.Errorf("expected: %v, but got: %v", n1, n2.GetPrev())
  }
}
