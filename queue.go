package main

// Queue is an interface that implementes queueing behavior
type Queue interface {
	Add(i interface{})
	Poll() interface{}
	Empty() bool
}
