package graph

import "testing"

func TestShortestPath(t *testing.T) {
	adjacencyList := make(map[string][]string)

	myFriends := []string{"Matias", "Paulina"}
	adjacencyList["Daniel"] = myFriends
	paulinasFriends := []string{"Matias", "Daniel"}
	adjacencyList["Paulina"] = paulinasFriends
	matiasFriends := []string{"Paulina", "Daniel", "Juan"}
	adjacencyList["Matias"] = matiasFriends
	juanFriends := []string{"Paulina", "Daniel", "Matias"}
	adjacencyList["Juan"] = juanFriends

	g := &Graph{
		adjacencyList: adjacencyList,
	}

	g.bfs()
}
