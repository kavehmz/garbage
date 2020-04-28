package main

import "fmt"

type graph struct {
	edges map[string]map[string]uint
	nodes map[string]bool
}

func (m *graph) addNode(name string) {
	if m.nodes == nil {
		m.nodes = make(map[string]bool)
	}
	m.nodes[name] = true
}

func (m *graph) addEdge(from, to string, cost uint) {
	if m.edges == nil {
		m.edges = make(map[string]map[string]uint)
	}
	if m.edges[from] == nil {
		m.edges[from] = make(map[string]uint)
	}
	m.edges[from][to] = cost
}

const inf = ^uint(0)

func (m *graph) dijkstra(from string) map[string]uint {
	costTable := make(map[string]uint)
	visited := make(map[string]bool)
	costTable[from] = 0

	for len(visited) < len(m.nodes) {
		nextNode := findSmallestUnvisited(costTable, visited)
		if nextNode == "" {
			return costTable
		}
		visited[nextNode] = true
		for toNode, toCost := range m.edges[nextNode] {
			if visited[toNode] {
				continue
			}
			if lastCost, found := costTable[toNode]; !found || costTable[nextNode]+toCost < lastCost {
				costTable[toNode] = costTable[nextNode] + toCost
			}
		}

	}

	return costTable
}

func findSmallestUnvisited(costTable map[string]uint, visited map[string]bool) string {
	smallestCost := inf
	smallestNode := ""
	for k, v := range costTable {
		if v < smallestCost && !visited[k] {
			smallestCost = v
			smallestNode = k
		}
	}
	return smallestNode
}

func main() {
	g := graph{}
	g.addNode("a")
	g.addNode("b")
	g.addNode("c")
	g.addNode("d")
	g.addNode("e")
	g.addNode("f")
	g.addNode("g")

	g.addEdge("a", "c", 2)
	g.addEdge("a", "b", 5)
	g.addEdge("c", "b", 1)
	g.addEdge("c", "d", 9)
	g.addEdge("b", "d", 4)
	g.addEdge("d", "e", 2)
	g.addEdge("d", "g", 30)
	g.addEdge("d", "f", 10)
	g.addEdge("f", "g", 1)

	fmt.Println(g.dijkstra("a"))
}
