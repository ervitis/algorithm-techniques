package main

import "fmt"

type graph struct {
	edges map[rune][]rune
}

func (g *graph) addEdge(from, to rune) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *graph) dfs(start, target rune, visited map[rune]bool) bool {
	if visited[start] {
		return false
	}
	visited[start] = true
	if start == target {
		return true
	}
	for _, edge := range g.edges[start] {
		if g.dfs(edge, target, visited) {
			return true
		}
	}
	return false
}

func newGraph() *graph {
	return &graph{make(map[rune][]rune)}
}

func main() {
	g := newGraph()
	g.addEdge('A', 'C')
	g.addEdge('C', 'E')
	g.addEdge('E', 'G')
	g.addEdge('G', 'F')
	g.addEdge('A', 'D')
	g.addEdge('D', 'F')
	fmt.Println(g.dfs('E', 'F', make(map[rune]bool)))
}
