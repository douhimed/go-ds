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
		sm.Put(fmt.Sprintf("user_%d", i), i)
	}

	if sm.Length() != 3 {
		t.Errorf("put : expected %d, actual %d", 3, sm.Length())
	}
}

func generateMap[K comparable](keys ...K) *SafeMap[K, int] {
	sm := NewMap[K, int]()

	for i := 0; i < len(keys); i++ {
		sm.Put(keys[i], i)
	}

	return sm
}

func TestGetOfMap(t *testing.T) {

	actual := []string{"med", "ahmed"}
	sm := generateMap(actual...)

	for i, value := range actual {

		if v, ok := sm.Get(value); !ok {
			t.Errorf("get : expected %d not found", i)
		} else if v != i {
			t.Errorf("get : expected %d actual %v", i, v)
		}

	}

	if _, ok := sm.Get("unknown"); ok {
		t.Errorf("get : expected %v not found", "unknown")
	}

}

func TestDelete(t *testing.T) {

	actual := []string{"med", "ahmed"}
	sm := generateMap(actual...)

	for i, v := range actual {

		if ok := sm.Delete(v); !ok {
			t.Errorf("delete : didn't delete the key %s", v)
		}

		if 1-i != sm.Length() {
			t.Errorf("delete : expected %d, actual %d", 1-i, sm.Length())
		}

	}

	if ok := sm.Delete("unknown"); ok {
		t.Error("delete : delete something not exist")
	}

}

/*********************************************
************ MAP TESTS : GO routines *********
**********************************************/

// RUN : go test -race

func TestWithGoRoutines(t *testing.T) {

	sm := NewMap[int, int]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			sm.Put(i, i*10)

			v, ok := sm.Get(i)

			if !ok {
				t.Errorf("go routine : didn't find %d", i)
			}

			if v != i*10 {
				t.Errorf("go routine : expected %d, actual %d", i*10, v)
			}

		}(i)
	}

}
