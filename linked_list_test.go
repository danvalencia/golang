package main

import (
	"fmt"
	"testing"
)

func TestCountList(t *testing.T) {
	list := LinkedList{}
	if list.Size() != 0 {
		t.Error("List should have a length of 0")
	}

	list.Append(1)
	if list.Size() != 1 {
		t.Error("List should have a length of 1")
	}

	list.Append(2)
	if list.Size() != 2 {
		t.Error("List should have a length of 2")
	}

	list.Remove()
	size := list.Size()
	if size != 1 {
		t.Errorf("List should have a length of 1 but was %v", size)
	}
}

func TestRemoveFromList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)

	element := list.Remove()
	if element != 2 {
		t.Errorf("Expected removed element to be 2 but was %v", element)
	}

	element = list.Remove()
	if element != 1 {
		t.Errorf("Expected removed element to be 1 but was %v", element)
	}
}

func TestAppendToList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println()
	fmt.Println(list.head, list.tail)
}

func TestPrependToList(t *testing.T) {
	list := LinkedList{}
	list.Prepend(1)
	list.Prepend(2)

	if list.head.value != 2 {
		t.Error("Expected head to be 2")
	}

	if list.tail.value != 1 {
		t.Error("Expected tail to be 1")
	}
}

func TestRemoveFirst(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	result := list.RemoveFirst()
	if result != 1 {
		t.Error("Expected element to be 1")
	}

	result = list.RemoveFirst()
	if result != 2 {
		t.Error("Expected element to be 2")
	}

	result = list.RemoveFirst()
	if result != 3 {
		t.Error("Expected element to be 3")
	}

	result = list.RemoveFirst()
	if result != -1 {
		t.Error("Expected element to be 3")
	}

	size := list.Size()
	if size != 0 {
		t.Errorf("Expected size to be 1 but was %v", size)
	}
}
