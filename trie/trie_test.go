package trie

import (
	"reflect"
	"sort"
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
		{name: "test char in incompleteWord", args: args{prefix: "a", dict: []string{"abc"}}, want: true},
		{name: "test chars in incompleteWord", args: args{prefix: "ab", dict: []string{"abc"}}, want: true},
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
		{name: "return 1 with prefix on a single incompleteWord dictionary", args: args{prefix: "a", dict: []string{"abc"}}, want: 1},
		{name: "adams's example Dog", args: args{prefix: "Dog", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 1},
		{name: "adams's example Ca", args: args{prefix: "Ca", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
		{name: "adams's example Zeg", args: args{prefix: "Zeg", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 0},
		{name: "adams's example Bea", args: args{prefix: "Bea", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
		{name: "adams's example Bear", args: args{prefix: "Bear", dict: []string{"Dog", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 1},
		{name: "incompleteWord and prefix", args: args{prefix: "Dog", dict: []string{"Dog", "Dogs", "Cat", "Bear", "Mouse", "Car", "Bean"}}, want: 2},
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

func TestTrie_FindFirstWith(t1 *testing.T) {
	type fields struct {
		Root *node
	}
	type args struct {
		char rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "empty trie returns empty string",
			fields: fields{Root: NewTrie().Root},
			args:   args{char: 'e'},
			want:   "",
		},
		{
			name:   "single incompleteWord with matching char",
			fields: fields{Root: NewTrie().Insert("bear").Root},
			args:   args{char: 'e'},
			want:   "bear",
		},
		{
			name: "example from prompt (ape, bear, tree, jeep)",
			fields: fields{Root: func() *node {
				t := NewTrie()
				t.Insert("ape")
				t.Insert("bear")
				t.Insert("tree")
				t.Insert("jeep")
				return t.Root
			}()},
			args: args{char: 'e'},
			want: "bear",
		},
		{
			name: "character appears earliest in multiple words, return lexicographically first",
			fields: fields{Root: func() *node {
				t := NewTrie()
				t.InsertMany([]string{"foo", "egg", "elk"})
				return t.Root
			}()},
			args: args{char: 'e'},
			want: "egg", // both egg and elk have 'e' at index 0, egg comes first
		},
		{
			name: "character not found in any incompleteWord",
			fields: fields{Root: func() *node {
				t := NewTrie()
				t.InsertMany([]string{"cat", "dog", "bird"})
				return t.Root
			}()},
			args: args{char: 'z'},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.fields.Root,
			}
			if got := t.FindFirstWith(tt.args.char); got != tt.want {
				t1.Errorf("FindFirstWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_FindAllWords(t1 *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "Empty trie",
			words: []string{},
			want:  nil,
		},
		{
			name:  "Single incompleteWord",
			words: []string{"hello"},
			want:  []string{"hello"},
		},
		{
			name:  "Multiple words",
			words: []string{"cat", "car", "card"},
			want:  []string{"car", "card", "cat"},
		},
		{
			name:  "Words with shared prefix",
			words: []string{"app", "apple", "application"},
			want:  []string{"app", "apple", "application"},
		},
		{
			name:  "Unrelated words",
			words: []string{"foo", "bar", "baz"},
			want:  []string{"bar", "baz", "foo"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			trie := NewTrie()
			trie.InsertMany(tt.words)
			got := trie.FindAllWords()
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("FindAllWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_FindAllLevels(t1 *testing.T) {
	type fields struct {
		Words []string
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		{
			name:   "empty trie",
			fields: fields{Words: []string{}},
			want:   [][]string{{""}},
		},
		{
			name:   "single incompleteWord",
			fields: fields{Words: []string{"a"}},
			want: [][]string{
				{""},
				{"a"},
			},
		},
		{
			name:   "multiple short words",
			fields: fields{Words: []string{"a", "b"}},
			want: [][]string{
				{""},
				{"a", "b"},
			},
		},
		{
			name:   "branching words",
			fields: fields{Words: []string{"car", "cat", "dog"}},
			want: [][]string{
				{""},
				{"c", "d"},
				{"a", "o"},
				{"r", "t", "g"},
			},
		},
		{
			name:   "deep single chain",
			fields: fields{Words: []string{"abcd"}},
			want: [][]string{
				{""},
				{"a"},
				{"b"},
				{"c"},
				{"d"},
			},
		},
		{
			name:   "mixed branching",
			fields: fields{Words: []string{"app", "apple", "arm"}},
			want: [][]string{
				{""},
				{"a"},
				{"p", "r"},
				{"p", "m"},
				{"l"},
				{"e"},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			trie := NewTrie()
			trie.InsertMany(tt.fields.Words)

			got := trie.FindAllLevels()

			normalizeLevels(got)
			normalizeLevels(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("FindAllLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func normalizeLevels(levels [][]string) {
	for i := range levels {
		sort.Strings(levels[i])
	}
}

func TestTrie_FindLongestWord(t1 *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "empty trie",
			words: []string{},
			want:  "",
		},
		{
			name:  "single incompleteWord",
			words: []string{"hello"},
			want:  "hello",
		},
		{
			name:  "multiple words different lengths",
			words: []string{"a", "ab", "abc", "abcd"},
			want:  "abcd",
		},
		{
			name:  "longest incompleteWord not last inserted",
			words: []string{"application", "app", "apple"},
			want:  "application",
		},
		{
			name:  "unrelated words",
			words: []string{"foo", "barbecue", "baz"},
			want:  "barbecue",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			trie := NewTrie()
			trie.InsertMany(tt.words)
			if got := trie.FindLongestWord(); got != tt.want {
				t1.Errorf("FindLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_FindLongestWord_SameLength(t1 *testing.T) {
	trie := NewTrie()
	trie.InsertMany([]string{"cat", "dog", "bat"})

	got := trie.FindLongestWord()
	valid := []string{"cat", "dog", "bat"}

	found := false
	for _, v := range valid {
		if got == v {
			found = true
			break
		}
	}
	if !found {
		t1.Errorf("FindLongestWord() = %v, want one of %v", got, valid)
	}
}
