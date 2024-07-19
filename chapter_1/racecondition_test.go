package chapter1

import "testing"

func TestRacecondition(t *testing.T) {
	out := racecondition()
	v, ok := out[1]
	if !ok {
		t.Error("expected key 1 being set")
	}
	if v {
		t.Error("expected key 1 being false")
	}
}
