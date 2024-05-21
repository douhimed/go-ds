package main

import (
	"testing"
)

func TestNewAndLengthOfEmptyList(t *testing.T) {

	list := New[string]()

	if list == nil {
		t.Error("it's nil")
	}

	if list.Length() != 0 {
		t.Error("must be empty")
	}

	list.Add("Med")
	list.Add("KHalid")

	if list.Length() != 2 {
		t.Errorf("length expected %d, but got %d", 2, list.Length())
	}
}
