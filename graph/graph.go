package graph

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
	queue := []int{source}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		if u == destination {
			return true
		}

		for _, v := range g.graph[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return false
}

func (g UndirectedGraph) IsFullyConnected() bool {
	if len(g.graph) == 1 {
		return true
	}
	visited := make([]bool, len(g.graph))
	visited[0] = true
	queue := []int{0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range g.graph[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	for i := range g.graph {
		if visited[i] == false {
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
