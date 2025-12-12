package graph

import "ds/queue"

type Edge struct {
	u, v int
}

func NewUndirectedGraph(n int, edges []Edge) *UndirectedGraph {
	g := make([][]int, n)
	for _, e := range edges {
		g[e.u] = append(g[e.u], e.v)
		g[e.v] = append(g[e.v], e.u)
	}
	return &UndirectedGraph{graph: g}
}

type UndirectedGraph struct {
	graph [][]int
}

func (g UndirectedGraph) IsThereAPathBetween(source, destination int) bool {
	if source == destination {
		return true
	}
	visited := make([]bool, len(g.graph))
	visited[source] = true
	q := []int{source}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		if u == destination {
			return true
		}

		for _, v := range g.graph[u] {
			if !visited[v] {
				visited[v] = true
				q = append(q, v)
			}
		}
	}
	return false
}

func (g UndirectedGraph) IsFullyConnected() bool {
	visited := make([]bool, len(g.graph))
	visited[0] = true
	q := queue.New[int]()
	q.Enqueue(0)
	for q.Len() > 0 {
		u := q.Dequeue()

		for _, v := range g.graph[u] {
			if !visited[v] {
				visited[v] = true
				q.Enqueue(v)
			}
		}
	}
	for _, v := range visited {
		if v == false {
			return false
		}
	}
	return true
}

func (g UndirectedGraph) IsBipartite() bool {
	if len(g.graph) == 1 {
		return false
	}
	colour := make([]int, len(g.graph))
	for i := range colour {
		colour[i] = -1
	}

	for i := range g.graph {
		if colour[i] == -1 {
			colour[i] = 0

			queue := []int{i}

			for len(queue) > 0 {
				u := queue[0]
				queue = queue[1:]

				for _, v := range g.graph[u] {
					if colour[v] == -1 {
						colour[v] = 1 - colour[u]
						queue = append(queue, v)
					} else if colour[v] == colour[u] {
						return false
					}
				}
			}
		}
	}
	return true
}

func (g UndirectedGraph) IsATree() bool {
	if len(g.graph) == 1 {
		return true
	}
	return g.IsFullyConnected() && !g.HasCycles()
}

func (g UndirectedGraph) HasCycles() bool {
	visited := make([]bool, len(g.graph))
	parent := make([]int, len(g.graph))
	for i := range parent {
		parent[i] = -1
	}
	for i := 0; i < len(g.graph); i++ {
		if !visited[i] {
			q := queue.New[int]()
			q.Enqueue(i)
			visited[i] = true

			for !q.Empty() {
				u := q.Dequeue()

				for _, v := range g.graph[u] {
					if !visited[v] {
						visited[v] = true
						parent[v] = u
						q.Enqueue(v)
					} else if v != parent[u] {
						return true
					}
				}
			}
		}
	}
	return false
}
