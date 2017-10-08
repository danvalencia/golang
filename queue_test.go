package main

import "testing"

func TestQueue(t *testing.T) {
	var queue Queue = &LinkedList{}

	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	if queue.Empty() {
		t.Error("Expected Queue to NOT be empty")
	}

	element := queue.Poll()
	if element != 1 {
		t.Errorf("Expected element to be 1 but was %v", element)
	}

	if queue.Empty() {
		t.Error("Expected Queue to NOT be empty")
	}

	element = queue.Poll()
	if element != 2 {
		t.Errorf("Expected element to be 2 but was %v", element)
	}

	if queue.Empty() {
		t.Error("Expected Queue to NOT be empty")
	}

	element = queue.Poll()
	if element != 3 {
		t.Errorf("Expected element to be 3 but was %v", element)
	}

	if !queue.Empty() {
		t.Error("Expected Queue to BE empty")
	}

}
