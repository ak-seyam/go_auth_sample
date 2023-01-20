package test

import "testing"

func TestAdding(t *testing.T) {
	got := 5 + 4
	want := 9
	if got != want {
		t.Error("assertion failed")
	}
}

func TestAddingFail(t *testing.T) {
	got := 5 + 4
	want := 10
	if got != want {
		t.Error("assertion failed")
	}
}
