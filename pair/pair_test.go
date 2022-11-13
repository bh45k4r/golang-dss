package pair

import "testing"

func TestPairStringInt(t *testing.T) {
	p := New("abcd", 1234)
	if k := p.Key(); k != "abcd" {
		t.Errorf("expected: \"abcd\", but got: %s", k)
	}
	if v := p.Value(); v != 1234 {
		t.Errorf("expected: 1234, but got: %d", v)
	}
}
