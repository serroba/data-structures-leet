package graph

import (
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
