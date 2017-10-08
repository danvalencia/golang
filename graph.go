package main

import (
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	adjacencyList map[string][]string
}

func (graph *Graph) bfs() {
	visiting := make(map[string]bool)

	var queue Queue = &LinkedList{}
	nodes := make([]string, len(graph.adjacencyList))

	i := 0
	for k := range graph.adjacencyList {
		nodes[i] = k
		i++
	}

	if len(nodes) > 0 {
		queue.Add(nodes[0])
	}

	for !queue.Empty() {
		currNode, ok := queue.Poll().(string)

		if ok {
			if !visiting[currNode] {
				fmt.Printf("Visiting %v\n", currNode)
				visiting[currNode] = true

				adjacentNodes := graph.adjacencyList[currNode]
				for _, v := range adjacentNodes {
					fmt.Printf("Adding %v to the queue\n", v)
					queue.Add(v)
				}
			} else {
				fmt.Printf("Already visited %v, skipping this one\n", currNode)
			}
		}
	}
}
