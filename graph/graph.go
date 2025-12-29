package graph

import (
	"math"

	"ds/queue"
)

type Edge struct {
	u, v int
}

type WeightedEdge struct {
	U, V, Weight int
}

type WeightedGraph struct {
	adj [][]struct{ to, weight int }
}

func NewWeightedGraph(n int, edges []WeightedEdge) *WeightedGraph {
	adj := make([][]struct{ to, weight int }, n)
	for _, e := range edges {
		adj[e.U] = append(adj[e.U], struct{ to, weight int }{e.V, e.Weight})
		adj[e.V] = append(adj[e.V], struct{ to, weight int }{e.U, e.Weight})
	}

	return &WeightedGraph{adj: adj}
}

// Dijkstra finds the shortest path from source to all other nodes.
// Returns the distance array where dist[i] is the shortest distance from source to node i.
// Unreachable nodes have distance math.MaxInt.
func (g *WeightedGraph) Dijkstra(source int) []int {
	n := len(g.adj)

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	dist[source] = 0

	pq := &queue.PriorityQueue{}
	pq.Push(0, source)

	for !pq.Empty() {
		item, _ := pq.Pop()
		d, u := item.Priority, item.Value

		// Skip if we've already found a shorter path
		if d > dist[u] {
			continue
		}

		for _, edge := range g.adj[u] {
			v, w := edge.to, edge.weight
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				pq.Push(dist[v], v)
			}
		}
	}

	return dist
}

// DijkstraPath finds the shortest path from source to destination.
// Returns the path as a slice of nodes and the total distance.
// If no path exists, returns nil and -1.
func (g *WeightedGraph) DijkstraPath(source, destination int) ([]int, int) {
	n := len(g.adj)
	dist := make([]int, n)
	prev := make([]int, n)

	for i := range dist {
		dist[i] = math.MaxInt
		prev[i] = -1
	}

	dist[source] = 0

	pq := &queue.PriorityQueue{}
	pq.Push(0, source)

	for !pq.Empty() {
		item, _ := pq.Pop()
		d, u := item.Priority, item.Value

		if u == destination {
			break
		}

		if d > dist[u] {
			continue
		}

		for _, edge := range g.adj[u] {
			v, w := edge.to, edge.weight
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
				pq.Push(dist[v], v)
			}
		}
	}

	if dist[destination] == math.MaxInt {
		return nil, -1
	}

	// Reconstruct path
	path := []int{}
	for at := destination; at != -1; at = prev[at] {
		path = append(path, at)
	}

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, dist[destination]
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
		if !v {
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
					switch colour[v] {
					case -1:
						colour[v] = 1 - colour[u]
						queue = append(queue, v)
					case colour[u]:
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

	for i := range len(g.graph) {
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
