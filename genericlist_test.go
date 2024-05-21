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

}

func TestAdd(t *testing.T) {

	list := New[string]()

	list.Add("Med")
	list.Add("KHalid")

	if list.Length() != 2 {
		t.Errorf("length expected %d, but got %d", 2, list.Length())
	}
}

func TestGet(t *testing.T) {

	list := New[string]()

	actual := []string{"med", "ahmed", "nabil"}

	for _, v := range actual {
		list.Add(v)
	}

	for i, v := range actual {
		if list.Get(i) != v {
			t.Errorf("Expected %s, actual %s", v, list.Get(i))
		}
	}
}
