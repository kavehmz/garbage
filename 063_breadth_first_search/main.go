package main

import (
	"fmt"
)

type graph struct {
	g map[string]map[string]bool
}

func (m *graph) addNode(name string) {
	if m.g == nil {
		m.g = make(map[string]map[string]bool)
	}
	m.g[name] = make(map[string]bool)
}

func (m *graph) addVert(from, to string) {
	m.g[from][to] = true
}

func connections(con map[string]bool) []string {
	q := []string{}
	for node := range con {
		q = append(q, node)
	}
	return q
}

func (m *graph) isConnected(from, to string) bool {
	queue := connections(m.g[from])
	visited := make(map[string]bool)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node == to {
			return true
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		queue = append(queue, connections(m.g[node])...)

	}
	return false
}

type queuedNode struct {
	name string
	from string
}

func (m *graph) connsQueue(from string) []queuedNode {
	q := []queuedNode{}
	con := m.g[from]
	for node := range con {
		q = append(q, queuedNode{name: node, from: from})
	}
	return q
}

func (m *graph) findPath(from, to string) []string {
	queue := m.connsQueue(from)
	visited := make(map[string]queuedNode)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if _, found := visited[node.name]; found {
			continue
		}
		visited[node.name] = queuedNode{name: node.name, from: node.from}
		if node.name == to {
			path := []string{}

			for node := visited[node.name]; node.name != ""; {
				path = append(path, node.name)
				node = visited[node.from]
			}
			path = append(path, from)
			return path
		}
		queue = append(queue, m.connsQueue(node.name)...)

	}
	return nil
}

func main() {
	/*
		graph = {}
		graph["you"] = ["alice", "bob", "claire"]
		graph["bob"] = ["anuj", "peggy"]
		graph["alice"] = ["peggy"]
		graph["claire"] = ["thom", "jonny"]
		graph["anuj"] = []
		graph["peggy"] = ["anuj"]
		graph["thom"] = []
		graph["jonny"] = []
	*/

	g := graph{}
	g.addNode("you")
	g.addVert("you", "alice")
	g.addVert("you", "bob")
	g.addVert("you", "claire")
	g.addNode("bob")
	g.addVert("bob", "anuj")
	g.addVert("bob", "peggy")
	g.addNode("alice")
	g.addVert("alice", "peggy")
	g.addNode("claire")
	g.addVert("claire", "thom")
	g.addVert("claire", "jonny")
	g.addNode("anuj")
	g.addNode("peggy")
	g.addVert("peggy", "anuj")
	g.addNode("thom")
	g.addNode("jonny")

	fmt.Println(g.isConnected("you", "thom"))
	fmt.Println(g.findPath("you", "thom"))
	fmt.Println(g.isConnected("bob", "claire"))
	fmt.Println(g.findPath("bob", "claire"))
	fmt.Println(g.isConnected("alice", "anuj"))
	fmt.Println(g.findPath("alice", "anuj"))

}
