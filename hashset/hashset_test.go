package hashset

import (
	"reflect"
	"sort"
	"testing"
)

func TestHashSetInt(t *testing.T) {
	hs := New[int]()
	hs.Add(1)
	hs.Add(3)
	hs.Add(5)
	if v := hs.Contains(4); v == true {
		t.Errorf("expected: false, but got: %v", v)
	}

	elements := make([]int, 0, 3)
	for v := range hs.Iterator() {
		elements = append(elements, v)
	}
	sort.Ints(elements)
	if !reflect.DeepEqual(elements, []int{1, 3, 5}) {
		t.Errorf("expected: [1,3,5], but got: %v", elements)
	}
}

func BenchmarkHashSetInt(b *testing.B) {
	hs := New[int]()
	for i := 0; i < b.N; i++ {
		hs.Add(i)
		if !hs.Contains(i) {
			b.Error("expected: true, but got: false")
		}
		hs.Remove(i)
	}
}
