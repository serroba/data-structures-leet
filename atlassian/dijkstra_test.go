package atlassian

import (
	"reflect"
	"testing"
)

func TestHasConnection(t *testing.T) {
	type args struct {
		g *Graph
		a rune
		b rune
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test that empty graph returns false", args: args{a: 1, b: 2}, want: false},
		{name: "test that graph without a", args: args{g: &Graph{
			g: map[rune]map[rune]int{'a': {'b': 1, 'c': 8}},
		}, a: 'a', b: 'b'}, want: false},
		{name: "test that graph", args: args{g: &Graph{
			g: map[rune]map[rune]int{'a': {'b': 1, 'c': 8}, 'b': {'a': 1, 'c': 6, 'd': 5}},
		}, a: 'a', b: 'b'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasConnection(tt.args.g, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("HasConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Dijkstra(t *testing.T) {
	tests := []struct {
		name   string
		graph  *Graph
		source rune
		want   map[rune]int
	}{
		{
			name: "simple path a-b-c",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 2},
				'b': {'a': 2, 'c': 3},
				'c': {'b': 3},
			}},
			source: 'a',
			want:   map[rune]int{'a': 0, 'b': 2, 'c': 5},
		},
		{
			name: "choose shorter path",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 1, 'c': 10},
				'b': {'a': 1, 'c': 2},
				'c': {'a': 10, 'b': 2},
			}},
			source: 'a',
			want:   map[rune]int{'a': 0, 'b': 1, 'c': 3}, // a->c via a->b->c = 3
		},
		{
			name: "disconnected node",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 5},
				'b': {'a': 5},
				'c': {}, // disconnected
			}},
			source: 'a',
			want:   map[rune]int{'a': 0, 'b': 5}, // c not reachable
		},
		{
			name: "classic example",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 4, 'c': 1},
				'b': {'a': 4, 'c': 2, 'd': 1},
				'c': {'a': 1, 'b': 2, 'd': 5},
				'd': {'b': 1, 'c': 5, 'e': 3},
				'e': {'d': 3},
			}},
			source: 'a',
			want:   map[rune]int{'a': 0, 'b': 3, 'c': 1, 'd': 4, 'e': 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.graph.Dijkstra(tt.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_DijkstraPath(t *testing.T) {
	tests := []struct {
		name        string
		graph       *Graph
		source      rune
		destination rune
		wantPath    []rune
		wantDist    int
	}{
		{
			name: "same source and destination",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 5},
				'b': {'a': 5},
			}},
			source:      'a',
			destination: 'a',
			wantPath:    []rune{'a'},
			wantDist:    0,
		},
		{
			name: "simple path",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 2},
				'b': {'a': 2, 'c': 3},
				'c': {'b': 3},
			}},
			source:      'a',
			destination: 'c',
			wantPath:    []rune{'a', 'b', 'c'},
			wantDist:    5,
		},
		{
			name: "path with shortcut",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 1, 'd': 100},
				'b': {'a': 1, 'c': 1},
				'c': {'b': 1, 'd': 1},
				'd': {'a': 100, 'c': 1},
			}},
			source:      'a',
			destination: 'd',
			wantPath:    []rune{'a', 'b', 'c', 'd'},
			wantDist:    3,
		},
		{
			name: "no path exists",
			graph: &Graph{g: map[rune]map[rune]int{
				'a': {'b': 1},
				'b': {'a': 1},
				'c': {'d': 1},
				'd': {'c': 1},
			}},
			source:      'a',
			destination: 'd',
			wantPath:    nil,
			wantDist:    -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotDist := tt.graph.DijkstraPath(tt.source, tt.destination)
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("DijkstraPath() path = %v, want %v", gotPath, tt.wantPath)
			}
			if gotDist != tt.wantDist {
				t.Errorf("DijkstraPath() dist = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
