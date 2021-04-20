package Graphs

import "fmt"

// 1 - > 2
// |
// v
// 3

// DAG
type Graph struct {
	v		int 
	adjList	map[interface{}][]interface{}
}

func New(V int) (Graph) {
	return Graph{
		adjList: make(map[interface{}][]interface{}),
		v: V,
	}
}

func (g *Graph) AddEdge(u interface{}, v interface{}) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *Graph) DFS(start interface{}) {
	visited := make(map[interface{}]bool, g.v)
	g._dfs(start, visited)
}

func (g* Graph) _dfs(u interface{}, visited map[interface{}]bool) {
	if visited[u] { 
		return
	}
	visited[u] = true 
	fmt.Printf("%v ", u)
	for i := 0; i < len(g.adjList[u]); i++ {
		v := g.adjList[u][i]
		g._dfs(v, visited)
	}
}

func (g *Graph) BFS(start interface{}) {
	visited := make(map[interface{}]bool, g.v)
	queue := make([]interface{}, 0)
	queue = append(queue, start) 

	for len(queue) > 0 {
		u := queue[0]
		fmt.Printf("%v ", u)
		queue = queue[1:]
		for i := 0; i < len(g.adjList[u]); i++ {
			v := g.adjList[u][i]
			if !visited[v] {
				visited[v] = true 
				queue = append(queue, v)
			}
		}
	}
	fmt.Println()
}

