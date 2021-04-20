package Graphs

import "fmt"

// 1 - > 2
// |
// v
// 3

// DAG
type Graph struct {
	v		int 
	adjList	map[interface{}]map[interface{}]int
}

func New(V int) (Graph) {
	return Graph{
		v: V,
	}
}

func (g *Graph) AddEdge(u interface{}, v interface{}) {
	g.adjList[u][v] = 1
}

func (g *Graph) BFS(start interface{}) {
	visited := make([]bool, g.v)
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

