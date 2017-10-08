package main

// Stack is a FIFO structure
type Stack interface {
	Push(i interface{})
	Pop() interface{}
}

// Push is the same as Append
func (l *LinkedList) Push(i interface{}) {
	l.Append(i)
}

// Pop is the same as Remove
func (l *LinkedList) Pop() interface{} {
	return l.Remove()
}
