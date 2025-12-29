package graph

import (
	"math"
	"reflect"
	"testing"
)

func TestUndirectedGraph_IsThereAPathBetween(t *testing.T) {
	type fields struct {
		n     int
		edges []Edge
	}
	type args struct {
		source      int
		destination int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "example 1", args: args{source: 0, destination: 2}, fields: fields{n: 3, edges: []Edge{{0, 1}, {1, 2}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph(tt.fields.n, tt.fields.edges)
			if got := g.IsThereAPathBetween(tt.args.source, tt.args.destination); !got {
				t.Errorf("IsThereAPathBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUndirectedGraph_CanReachAllNodes(t *testing.T) {
	type fields struct {
		n     int
		edges []Edge
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "check on single node", fields: fields{n: 1}, want: true},
		{name: "check on bigger graph", fields: fields{n: 3, edges: []Edge{{0, 1}, {1, 2}}}, want: true},
		{name: "check on bigger graph", fields: fields{n: 4, edges: []Edge{{0, 1}, {2, 3}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph(tt.fields.n, tt.fields.edges)
			if got := g.IsFullyConnected(); got != tt.want {
				t.Errorf("IsFullyConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUndirectedGraph_IsBipartite(t *testing.T) {
	type fields struct {
		n     int
		edges []Edge
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "single node graph",
			fields: fields{n: 1, edges: nil},
			want:   false,
		},
		{
			name: "simple bipartite square",
			fields: fields{
				n:     4,
				edges: []Edge{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			},
			want: true,
		},
		{
			name: "triangle (odd cycle) not bipartite",
			fields: fields{
				n:     3,
				edges: []Edge{{0, 1}, {1, 2}, {2, 0}},
			},
			want: false,
		},
		{
			name: "two disconnected components, both bipartite",
			fields: fields{
				n:     5,
				edges: []Edge{{0, 1}, {2, 3}},
			},
			want: true,
		},
		{
			name: "disconnected but one component not bipartite",
			fields: fields{
				n:     6,
				edges: []Edge{{0, 1}, {1, 2}, {2, 0}, {3, 4}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph(tt.fields.n, tt.fields.edges)
			if got := g.IsBipartite(); got != tt.want {
				t.Errorf("IsBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUndirectedGraph_IsATree(t *testing.T) {
	type fields struct {
		n     int
		edges []Edge
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "single node no edges is a tree",
			fields: fields{n: 1, edges: nil},
			want:   true,
		},
		{
			name:   "simple line tree",
			fields: fields{n: 3, edges: []Edge{{0, 1}, {1, 2}}},
			want:   true,
		},
		{
			name:   "star shaped tree",
			fields: fields{n: 4, edges: []Edge{{0, 1}, {0, 2}, {0, 3}}},
			want:   true,
		},
		{
			name:   "cycle of three is not a tree",
			fields: fields{n: 3, edges: []Edge{{0, 1}, {1, 2}, {2, 0}}},
			want:   false,
		},
		{
			name:   "disconnected graph is not a tree",
			fields: fields{n: 4, edges: []Edge{{0, 1}, {2, 3}}},
			want:   false,
		},
		{
			name: "self-loop is not a tree",
			fields: fields{
				n:     3,
				edges: []Edge{{0, 0}, {0, 1}, {1, 2}},
			},
			want: false,
		},
		{
			name: "tree with many nodes (chain of 10)",
			fields: fields{
				n:     10,
				edges: []Edge{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}},
			},
			want: true,
		},
		{
			name: "graph with multiple cycles",
			fields: fields{
				n: 7,
				edges: []Edge{
					{0, 1}, {1, 2}, {2, 0}, // first cycle
					{3, 4}, {4, 5}, {5, 3}, // second cycle
					{2, 3}, // connecting two cyclic components
				},
			},
			want: false,
		},
		{
			name: "graph almost a tree but with one extra edge forming a deep cycle",
			fields: fields{
				n: 6,
				edges: []Edge{
					{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, // linear chain
					{5, 1}, // extra back-edge creates a cycle
				},
			},
			want: false,
		},
		{
			name: "large star tree",
			fields: fields{
				n: 8,
				edges: []Edge{
					{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7},
				},
			},
			want: true,
		},
		{
			name: "two trees connected by an extra edge making a cycle",
			fields: fields{
				n: 6,
				edges: []Edge{
					{0, 1}, {1, 2}, // tree component A
					{3, 4}, {4, 5}, // tree component B
					{2, 3}, // bridging edge
					{5, 0}, // extra edge creating a cycle across components
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph(tt.fields.n, tt.fields.edges)
			if got := g.IsATree(); got != tt.want {
				t.Errorf("IsATree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeightedGraph_Dijkstra(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  []WeightedEdge
		source int
		want   []int
	}{
		{
			name:   "single node",
			n:      1,
			edges:  nil,
			source: 0,
			want:   []int{0},
		},
		{
			name: "simple path",
			n:    3,
			edges: []WeightedEdge{
				{0, 1, 2},
				{1, 2, 3},
			},
			source: 0,
			want:   []int{0, 2, 5},
		},
		{
			name: "choose shorter path",
			n:    4,
			edges: []WeightedEdge{
				{0, 1, 1},
				{1, 3, 1},
				{0, 2, 10},
				{2, 3, 1},
			},
			source: 0,
			want:   []int{0, 1, 3, 2}, // 0->2 via 0->1->3->2 is 3, shorter than direct 10
		},
		{
			name: "disconnected node",
			n:    4,
			edges: []WeightedEdge{
				{0, 1, 5},
				{1, 2, 3},
			},
			source: 0,
			want:   []int{0, 5, 8, math.MaxInt},
		},
		{
			name: "classic example graph",
			n:    5,
			edges: []WeightedEdge{
				{0, 1, 4},
				{0, 2, 1},
				{2, 1, 2},
				{1, 3, 1},
				{2, 3, 5},
				{3, 4, 3},
			},
			source: 0,
			want:   []int{0, 3, 1, 4, 7}, // 0->1 via 0->2->1 = 3, 0->3 via 0->2->1->3 = 4
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWeightedGraph(tt.n, tt.edges)
			got := g.Dijkstra(tt.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeightedGraph_DijkstraPath(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       []WeightedEdge
		source      int
		destination int
		wantPath    []int
		wantDist    int
	}{
		{
			name:        "same source and destination",
			n:           3,
			edges:       []WeightedEdge{{0, 1, 5}, {1, 2, 3}},
			source:      0,
			destination: 0,
			wantPath:    []int{0},
			wantDist:    0,
		},
		{
			name:        "simple path",
			n:           3,
			edges:       []WeightedEdge{{0, 1, 2}, {1, 2, 3}},
			source:      0,
			destination: 2,
			wantPath:    []int{0, 1, 2},
			wantDist:    5,
		},
		{
			name: "path with shortcut",
			n:    4,
			edges: []WeightedEdge{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 1},
				{0, 3, 10},
			},
			source:      0,
			destination: 3,
			wantPath:    []int{0, 1, 2, 3},
			wantDist:    3,
		},
		{
			name: "no path exists",
			n:    4,
			edges: []WeightedEdge{
				{0, 1, 1},
				{2, 3, 1},
			},
			source:      0,
			destination: 3,
			wantPath:    nil,
			wantDist:    -1,
		},
		{
			name: "complex graph find optimal path",
			n:    6,
			edges: []WeightedEdge{
				{0, 1, 7},
				{0, 2, 9},
				{0, 5, 14},
				{1, 2, 10},
				{1, 3, 15},
				{2, 3, 11},
				{2, 5, 2},
				{3, 4, 6},
				{4, 5, 9},
			},
			source:      0,
			destination: 4,
			wantPath:    []int{0, 2, 5, 4},
			wantDist:    20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWeightedGraph(tt.n, tt.edges)
			gotPath, gotDist := g.DijkstraPath(tt.source, tt.destination)
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("DijkstraPath() path = %v, want %v", gotPath, tt.wantPath)
			}
			if gotDist != tt.wantDist {
				t.Errorf("DijkstraPath() dist = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
