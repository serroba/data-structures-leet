package trie

import (
	"fmt"
	"sort"
	"strings"

	"ds/queue"
	"ds/stack"
)

type node struct {
	isEnd    bool
	children map[rune]*node
}

func (n *node) hasChildren() bool {
	return len(n.children) > 0
}

func (n *node) sortedChildrenKeys() []rune {
	keys := make([]rune, 0, len(n.children))
	for k := range n.children {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
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

func (t *Trie) SearchFirstWith(pattern string) []string {
	return []string{}
}

func (t *Trie) FindFirstWith(target rune) string {
	if t.Root == nil {
		return ""
	}
	type queueItem struct {
		node           *node
		incompleteWord string
	}
	q := queue.New[queueItem]()
	q.Enqueue(queueItem{node: t.Root, incompleteWord: ""})

	for !q.Empty() {
		item := q.Dequeue()
		for _, ch := range item.node.sortedChildrenKeys() {
			child := item.node.children[ch]
			nextWord := item.incompleteWord + string(ch)
			if ch == target {
				return nextWord + (&Trie{Root: child}).FindFirstSortedFullWord()
			}
			q.Enqueue(queueItem{node: child, incompleteWord: nextWord})
		}
	}
	return ""
}

func (t *Trie) FindFirstSortedFullWord() string {
	if t.Root == nil {
		return ""
	}
	type item struct {
		node  *node
		depth int
		word  string
	}
	s := stack.New(item{t.Root, 0, ""})
	for !s.Empty() {
		i, _ := s.Pop()
		if i.node.isEnd {
			return i.word
		}
		keys := i.node.sortedChildrenKeys()
		for j := len(keys) - 1; j >= 0; j-- {
			s.Push(item{i.node.children[keys[j]], i.depth + 1, i.word + string(keys[j])})
		}
	}
	return ""
}

func (t *Trie) FindAllWords() []string {
	var words []string
	type item struct {
		node *node
		word string
	}
	q := queue.New[item]()
	q.Enqueue(item{node: t.Root, word: ""})
	for !q.Empty() {
		i := q.Dequeue()
		for ch, n := range i.node.children {
			nextWord := i.word + string(ch)
			if n.isEnd {
				words = append(words, nextWord)
			}
			q.Enqueue(item{node: n, word: nextWord})
		}

	}
	return words
}

func findEarliestFullWord(node *node, prefix string) string {
	if node.isEnd {
		return prefix
	}

	for _, ch := range node.sortedChildrenKeys() {
		child := node.children[ch]
		result := findEarliestFullWord(child, prefix+string(ch))
		if result != "" {
			return result
		}
	}
	return ""
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
	//if !t.SearchWord(incompleteWord) {
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

func (t *Trie) FindAllLevels() [][]string {
	if t.Root == nil {
		return nil
	}
	current := t.Root
	q := queue.New[*node]()
	q.Enqueue(current)
	levels := [][]string{{""}}
	for !q.Empty() {
		var chars []string
		l := q.Len()
		for i := 0; i < l; i++ {
			n := q.Dequeue()
			for ch, val := range n.children {
				chars = append(chars, string(ch))
				q.Enqueue(val)
			}
		}
		if chars != nil {
			levels = append(levels, chars)
		}
	}
	return levels
}

func (t *Trie) FindLongestWord() string {
	if t.Root == nil {
		return ""
	}
	var longestWord string
	current := t.Root
	type item struct {
		node *node
		word string
	}
	q := queue.New(item{node: current, word: ""})
	for !q.Empty() {
		n := q.Dequeue()
		for ch, val := range n.node.children {
			nextWord := n.word + string(ch)
			if val.isEnd {
				if len(nextWord) > len(longestWord) {
					longestWord = nextWord
				}
			}
			q.Enqueue(item{node: val, word: nextWord})
		}
	}
	return longestWord
}

func (t *Trie) HasWordOf(length int) bool {
	if t.Root == nil {
		return false
	}
	type item struct {
		node  *node
		depth int
	}
	s := stack.New(item{node: t.Root, depth: 0})
	for !s.Empty() {
		i, _ := s.Pop()
		if length == i.depth {
			if i.node.isEnd {
				return true
			}
			continue
		}
		for _, child := range i.node.children {
			s.Push(item{child, i.depth + 1})
		}
	}
	return false
}
