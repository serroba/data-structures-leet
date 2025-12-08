package trie

import (
	"fmt"
	"sort"
	"strings"
)

type node struct {
	isEnd    bool
	children map[rune]*node
}

func (n *node) hasChildren() bool {
	return len(n.children) > 0
}

func (n *node) countWords() int {
	if n == nil {
		return 0
	}
	count := 0
	current := n
	if current.isEnd {
		count++
	}
	for _, ch := range current.children {
		count += ch.countWords()
	}
	return count
}

type Trie struct {
	Root *node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &node{children: make(map[rune]*node)},
	}
}

func (t *Trie) Insert(word string) *Trie {
	current := t.Root
	for _, ch := range word {
		if _, ok := current.children[ch]; !ok {
			current.children[ch] = &node{children: make(map[rune]*node)}
		}
		current = current.children[ch]
	}
	current.isEnd = true
	return t
}

func (t *Trie) InsertMany(words []string) {
	for _, word := range words {
		t.Insert(word)
	}
}

func (t *Trie) SearchPrefix(prefix string) bool {
	current := t.Root
	for _, ch := range prefix {
		if _, ok := current.children[ch]; !ok {
			return false
		}
		current = current.children[ch]
	}
	return true
}

func (t *Trie) SearchWord(word string) bool {
	if word == "" {
		return true
	}
	current := t.Root
	for _, ch := range word {
		if _, ok := current.children[ch]; !ok {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

func (t *Trie) SearchWith(pattern string) []string {
	return []string{}
}

func (t *Trie) CountWordsWith(prefix string) int {
	current := t.Root
	for _, ch := range prefix {
		if _, ok := current.children[ch]; !ok {
			return 0
		}
		current = current.children[ch]
	}

	return current.countWords()
}

func (t *Trie) Delete(word string) {
	//if !t.SearchWord(word) {
	//	return
	//}

	current := t.Root
	safeNode := t.Root
	var safeKey rune

	for _, ch := range word {
		if len(current.children) > 1 || current.isEnd {
			safeNode = current
			safeKey = ch
		}
		current = current.children[ch]
	}

	if current.hasChildren() {
		current.isEnd = false
	} else {
		delete(safeNode.children, safeKey)
	}
}

func (t *Trie) Equal(other *Trie) bool {
	if t == other {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	return equalNode(t.Root, other.Root)
}

func equalNode(a, b *node) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.isEnd != b.isEnd {
		return false
	}
	if len(a.children) != len(b.children) {
		return false
	}
	for k, ac := range a.children {
		bc, ok := b.children[k]
		if !ok || !equalNode(ac, bc) {
			return false
		}
	}
	return true
}

func (t *Trie) String() string {
	return renderNode(t.Root, 0, '∅')
}

func renderNode(n *node, depth int, ch rune) string {
	indent := ""
	if depth > 0 {
		indent = strings.Repeat("  ", depth)
	}

	end := ""
	if n.isEnd {
		end = " (end)"
	}

	line := fmt.Sprintf("%s[%c]%s\n", indent, ch, end)

	if len(n.children) == 0 {
		return line
	}

	keys := make([]rune, 0, len(n.children))
	for k := range n.children {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	result := line
	for _, k := range keys {
		result += renderNode(n.children[k], depth+1, k)
	}
	return result
}
