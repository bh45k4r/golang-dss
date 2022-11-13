package hashset


import "testing"


func TestHashSetInt(t *testing.T) {
  hs := New[int]()
  hs.Add(1)
  hs.Add(3)
  hs.Add(5)
  if v := hs.Contains(4); v == true {
    t.Errorf("expected: false, but got: %v", v)
  }
}
