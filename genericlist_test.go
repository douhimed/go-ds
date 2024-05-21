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

func TestRemoveUnvalideIndex(t *testing.T) {

	list := New[string]()

	unvalideIndexes := []int{-1, 99}

	for _, value := range unvalideIndexes {
		err := list.Remove(value)
		if err == nil {
			t.Errorf("out range check is KO")
		}
	}
}

func TestRemove(t *testing.T) {

	list := New[string]()

	actual := []string{"med", "ahmed", "nabil"}

	for _, v := range actual {
		list.Add(v)
	}

	for i := 2; i >= 0; i-- {
		err := list.Remove(i)

		if err != nil {
			t.Error(err.Error())
		}

		if list.Length() != i {
			t.Errorf("Expected %d, actual %d", i, list.Length())
		}
	}

}
