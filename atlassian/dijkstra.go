package atlassian

import (
	"math"

	"ds/queue"
)

type Graph struct {
	g map[rune]map[rune]int
}

// Dijkstra finds the shortest distances from source to all reachable nodes.
// Returns a map where dist[node] is the shortest distance from source to node.
// Unreachable nodes will not be present in the map.
func (g *Graph) Dijkstra(source rune) map[rune]int {
	dist := make(map[rune]int)
	for node := range g.g {
		dist[node] = math.MaxInt
	}

	dist[source] = 0

	pq := &queue.PriorityQueue{}
	pq.Push(0, int(source))

	for !pq.Empty() {
		item, _ := pq.Pop()
		d, u := item.Priority, rune(item.Value)

		if d > dist[u] {
			continue
		}

		for v, weight := range g.g[u] {
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				pq.Push(dist[v], int(v))
			}
		}
	}

	// Remove unreachable nodes
	for node, d := range dist {
		if d == math.MaxInt {
			delete(dist, node)
		}
	}

	return dist
}

// DijkstraPath finds the shortest path from source to destination.
// Returns the path as a slice of runes and the total distance.
// If no path exists, returns nil and -1.
func (g *Graph) DijkstraPath(source, destination rune) ([]rune, int) {
	dist := make(map[rune]int)
	prev := make(map[rune]rune)
	hasPrev := make(map[rune]bool)

	for node := range g.g {
		dist[node] = math.MaxInt
	}

	dist[source] = 0

	pq := &queue.PriorityQueue{}
	pq.Push(0, int(source))

	for !pq.Empty() {
		item, _ := pq.Pop()
		d, u := item.Priority, rune(item.Value)

		if u == destination {
			break
		}

		if d > dist[u] {
			continue
		}

		for v, weight := range g.g[u] {
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				prev[v] = u
				hasPrev[v] = true
				pq.Push(dist[v], int(v))
			}
		}
	}

	if dist[destination] == math.MaxInt {
		return nil, -1
	}

	// Reconstruct path
	path := []rune{}
	for at := destination; ; at = prev[at] {
		path = append(path, at)
		if !hasPrev[at] {
			break
		}
	}

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, dist[destination]
}

func HasConnection(g *Graph, a, b rune) bool {
	if g == nil {
		return false
	}

	if _, ok := g.g[a]; !ok {
		return false
	}

	if _, ok := g.g[b]; !ok {
		return false
	}

	var dfs func(g *Graph, a, b rune) bool

	dfs = func(g *Graph, a, b rune) bool {
		if a == b {
			return true
		}

		for neighbor := range g.g[a] {
			return dfs(g, neighbor, b)
		}

		return false
	}

	return dfs(g, a, b)
}
