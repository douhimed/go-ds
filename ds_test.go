package main

import (
	"fmt"
	"testing"
)

func create[T comparable](values ...T) *GenericList[T] {
	list := NewList[T]()

	for _, v := range values {
		list.Add(v)
	}

	return list
}

func TestNewAndLengthOfEmptyList(t *testing.T) {

	list := NewList[string]()

	if list == nil {
		t.Error("new : it's nil")
	}

	if list.Length() != 0 {
		t.Error("length : must be empty")
	}

}

func TestAdd(t *testing.T) {

	list := create("Med", "KHalid")

	if list.Length() != 2 {
		t.Errorf("add : length expected %d, but got %d", 2, list.Length())
	}
}

func TestGet(t *testing.T) {

	actual := []string{"Med", "ahmed", "nabil"}
	list := create(actual...)

	for i, v := range actual {
		if list.Get(i) != v {
			t.Errorf("add : expected %s, actual %s", v, list.Get(i))
		}
	}
}

func TestRemoveUnvalideIndex(t *testing.T) {

	list := NewList[string]()

	unvalideIndexes := []int{-1, 99}

	for _, value := range unvalideIndexes {
		err := list.Remove(value)
		if err == nil {
			t.Errorf("remove : out range check is KO")
		}
	}
}

func TestRemove(t *testing.T) {

	actual := []string{"Med", "ahmed", "nabil"}
	list := create(actual...)

	for i := 2; i >= 0; i-- {
		err := list.Remove(i)

		if err != nil {
			t.Error(err.Error())
		}

		if list.Length() != i {
			t.Errorf("remove : xpected %d, actual %d", i, list.Length())
		}
	}

}

func TestFirstIndex(t *testing.T) {

	actual := []string{"med", "ahmed", "ahmed"}
	list := create(actual...)

	j, ok := list.FirstIndex("med")
	if !ok || j != 0 {
		t.Errorf("firstIndex : expected %d actual %d", 0, j)
	}

	j, ok = list.FirstIndex("ahmed")
	if !ok || j != 1 {
		t.Errorf("firstIndex : expected %d actual %d", 1, j)
	}

	j, ok = list.FirstIndex("unknown")
	if ok || j != -1 {
		t.Errorf("firstIndex : expected %d actual %d", 0, j)
	}
}

func TestRemoveByValue(t *testing.T) {

	actual := []string{"Med", "ahmed", "ahmed"}
	list := create(actual...)

	for i := 2; i < len(actual); i++ {

		j, ok := list.RemoveByValue(actual[i])

		if j < 0 || !ok || list.Length() != i {
			t.Errorf("removeByValue : expected %d actual %d ", i, j)
		}

	}
}

/*********************************************
****************** MAP TESTS *****************
**********************************************/

func TestNewAndLengthSafeMap(t *testing.T) {

	sm := NewMap[string, int]()

	if sm == nil {
		t.Error("new : map is nil")
	}

	if sm.Length() != 0 {
		t.Errorf("length : expected %d, actual %d", 0, sm.Length())
	}
}

func TestPut(t *testing.T) {

	sm := NewMap[string, int]()

	for i := 0; i < 3; i++ {
		sm.put(fmt.Sprintf("user_%d", i), i)
	}

	if sm.Length() != 3 {
		t.Errorf("put : expected %d, actual %d", 3, sm.Length())
	}
}
