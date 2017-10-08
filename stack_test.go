package main

import "testing"

func TestPushAndPop(t *testing.T) {
	var stack Stack = &LinkedList{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	n := stack.Pop()
	if n != 3 {
		t.Errorf("Expected n to be 3 but was %v", n)
	}

	n = stack.Pop()
	if n != 2 {
		t.Errorf("Expected n to be 2 but was %v", n)
	}

	n = stack.Pop()
	if n != 1 {
		t.Errorf("Expected n to be 1 but was %v", n)
	}

	n = stack.Pop()
	if n != nil {
		t.Errorf("Expected n to be nil but was %v", n)
	}

	n = stack.Pop()
	if n != nil {
		t.Errorf("Expected n to be nil but was %v", n)
	}
}
