package trie

import (
	"reflect"
	"testing"
)

func TestTrie_SearchPrefix(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		prefix string
		dict   []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "test empty string on empty trie", args: args{prefix: ""}, want: true},
		{name: "test empty string on non-empty trie", args: args{prefix: "", dict: []string{"a", "b"}}, want: true},
		{name: "test single char", args: args{prefix: "a", dict: []string{"a"}}, want: true},
		{name: "test char in word", args: args{prefix: "a", dict: []string{"abc"}}, want: true},
		{name: "test chars in word", args: args{prefix: "ab", dict: []string{"abc"}}, want: true},
		{name: "test chars in many words", args: args{prefix: "ab", dict: []string{"abc", "abd", "ab", "car"}}, want: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTrie()
			t.InsertMany(tt.args.dict)
			if got := t.SearchPrefix(tt.args.prefix); got != tt.want {
				t1.Errorf("SearchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_SearchWord(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		word string
		dict []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "test empty string on empty trie", args: args{word: ""}, want: true},
		{name: "test empty string on non-empty trie", args: args{word: "", dict: []string{"a", "b"}}, want: true},
		{name: "test single char", args: args{word: "a", dict: []string{"a"}}, want: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTrie()
			t.InsertMany(tt.args.dict)
			if got := t.SearchWord(tt.args.word); got != tt.want {
				t1.Errorf("SearchWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_CountWordsWith(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		prefix string
		dict   []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "return 1 with no prefix", args: args{prefix: ""}, want: 0},
		{name: "return 1 with prefix on a single word dictionary", args: args{prefix: "a", dict: []string{"abc"}}, want: 1},
		{name: "adams's example Dog", args: args{prefix: "Dog", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 1},
		{name: "adams's example Ca", args: args{prefix: "Ca", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
		{name: "adams's example Zeg", args: args{prefix: "Zeg", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 0},
		{name: "adams's example Bea", args: args{prefix: "Bea", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
		{name: "adams's example Bear", args: args{prefix: "Bear", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 1},
		{name: "word and prefix", args: args{prefix: "Dog", dict: []string{"Dog", "Dogs", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
		{name: "count all words", args: args{prefix: "", dict: []string{"Dog", "Dogs", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 7},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTrie()
			t.InsertMany(tt.args.dict)
			if got := t.CountWordsWith(tt.args.prefix); got != tt.want {
				t1.Errorf("CountWordsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Delete(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		word string
		dict []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Trie
	}{
		{name: "delete empty", args: args{word: "not in dict"}, want: NewTrie()},
		{name: "delete not in dict", args: args{word: "not in dict", dict: []string{"some", "words"}}, want: NewTrie().Insert("some").Insert("words")},
		{name: "delete cat", args: args{word: "cat", dict: []string{"car", "cat", "caterpillar"}}, want: NewTrie().Insert("car").Insert("caterpillar")},
		{name: "delete cat", args: args{word: "cars", dict: []string{"cars", "carpool", "caterpillar"}}, want: NewTrie().Insert("carpool").Insert("caterpillar")},
		{name: "delete caterpillar", args: args{word: "caterpillar", dict: []string{"car", "cat", "caterpillar"}}, want: NewTrie().Insert("car").Insert("cat")},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTrie()
			t.InsertMany(tt.args.dict)
			t.Delete(tt.args.word)
			if got := t; !got.Equal(tt.want) {
				t1.Errorf("Trie = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestTrie_SearchWith(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		wildcard string
		dict     []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{name: "Search with c?", args: args{wildcard: "c?", dict: []string{"cars", "carpool", "caterpillar"}}, want: []string{"cars", "carpool", "caterpillar"}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTrie()
			t.InsertMany(tt.args.dict)
			if got := t.SearchWith(tt.args.wildcard); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SearchWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
