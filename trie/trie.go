package trie

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"ds/queue"
	"ds/stack"
)

type node struct {
	isEnd    bool
	children map[rune]*node
	parent   *node
	val      rune
}

func (n *node) hasChildren() bool {
	return len(n.children) > 0
}

func (n *node) sortedChildrenKeys() []rune {
	keys := make([]rune, 0, len(n.children))
	for k := range n.children {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (n *node) countWords() int {
	if n == nil {
		return 0
	}
	count := 0
	q := queue.New(n)
	for !q.Empty() {
		i := q.Dequeue()
		if i.isEnd {
			count++
		}
		for _, child := range i.children {
			q.Enqueue(child)
		}
	}
	return count
}

func (n *node) getWordSoFar() string {
	if n == nil {
		return ""
	}
	current := n
	wordSoFar := ""
	for current != nil {
		wordSoFar = string(current.val) + wordSoFar
		current = current.parent
	}
	return wordSoFar
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
	var parent *node
	for _, ch := range word {
		if _, ok := current.children[ch]; !ok {
			current.children[ch] = &node{children: make(map[rune]*node), val: ch}

			if parent != nil {
				current.children[ch].parent = parent
			}
		}
		current = current.children[ch]
		if current.val != 0 {
			current.parent = parent
		}
		parent = current
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

// SearchWith finds all words matching a pattern with '?' as single-character wildcard.
// It tracks multiple positions in the trie simultaneously, advancing through the pattern.
//
// Example: pattern "?at" on trie with ["cat", "bat", "rat", "car"]
//
//	Trie structure:
//	        root
//	       / | \
//	      c  b  r
//	     /|  |   \
//	    a o  a    a
//	   /| |  |     \
//	  t r t  t      t
//
//	Step 1: patternChar = '?'
//	  frontier: [root]
//	  '?' matches any char → expand to ALL children
//	  nextFrontier: [(c,"c"), (b,"b"), (r,"r")]
//
//	Step 2: patternChar = 'a'
//	  frontier: [(c,"c"), (b,"b"), (r,"r")]
//	  'a' is literal → only keep nodes with 'a' child
//	  nextFrontier: [(ca,"ca"), (ba,"ba"), (ra,"ra")]
//
//	Step 3: patternChar = 't'
//	  frontier: [(ca,"ca"), (ba,"ba"), (ra,"ra")]
//	  't' is literal → only keep nodes with 't' child
//	  nextFrontier: [(cat,"cat"), (bat,"bat"), (rat,"rat")]
//
//	Collect: all three nodes have isEnd=true
//	Result: ["cat", "bat", "rat"]
func (t *Trie) SearchWith(pattern string) []string {
	if t.Root == nil || t.Root.children == nil {
		return nil
	}
	type candidate struct {
		node      *node
		wordSoFar string
	}

	// frontier holds all current positions we're tracking in parallel
	frontier := []candidate{{node: t.Root}}

	for _, patternChar := range pattern {
		var nextFrontier []candidate

		for _, current := range frontier {
			if patternChar == '?' {
				// wildcard: branch out to ALL children
				for childChar, childNode := range current.node.children {
					nextFrontier = append(nextFrontier, candidate{
						node:      childNode,
						wordSoFar: current.wordSoFar + string(childChar),
					})
				}
			} else {
				// literal: follow only the matching child (if exists)
				if childNode, ok := current.node.children[patternChar]; ok {
					nextFrontier = append(nextFrontier, candidate{
						node:      childNode,
						wordSoFar: current.wordSoFar + string(patternChar),
					})
				}
			}
		}

		frontier = nextFrontier
	}

	// collect results: only nodes that mark end of a complete word
	var matches []string
	for _, current := range frontier {
		if current.node.isEnd {
			matches = append(matches, current.wordSoFar)
		}
	}
	return matches
}

func (t *Trie) FindFirstWordMatching(pattern string) string {
	if t.Root == nil {
		return ""
	}
	if pattern == "" {
		return ""
	}
	current := t.Root
	for _, ch := range pattern {
		if _, ok := current.children[ch]; !ok && ch != '?' {
			return ""
		}
		if ch == '?' {
			for key, _ := range current.children {
				current = current.children[key]
				break
			}
		} else {
			current = current.children[ch]
		}
	}
	if !current.isEnd {
		return ""
	}
	return current.getWordSoFar()
}

func GetWordSoFar(node *node) string {
	var wordSoFar = string(node.val)
	for node.parent != nil {
		node = node.parent
		wordSoFar = string(node.val) + wordSoFar
	}
	return wordSoFar
}

// patern c?t
// pattern t?y  word = turkey
func (n *node) FindAllMatching(pattern string) []string {
	var matches []string
	if pattern == "" {
		if n.isEnd {
			return []string{n.getWordSoFar()}
		}
		return nil
	}
	if !n.hasChildren() {
		return nil
	}

	var matchingChildren map[rune]*node
	ch := []rune(pattern)[0]
	if _, ok := n.children[ch]; !ok && ch != '?' {
		return nil
	}
	if ch == '?' {
		// Find all words under all children
		matchingChildren = n.children
	} else {
		// Find all words under a single matching child
		matchingChildren = map[rune]*node{ch: n.children[ch]}
	}

	for _, child := range matchingChildren {
		matches = append(matches, child.FindAllMatching(pattern[1:])...)
	}
	return matches
}

// pattern t*y  word = turkey
func (n *node) FindStarMatching(pattern string) []string {
	if pattern == "" {
		if n.isEnd {
			// Pattern fully consumed. This is a match!
			return []string{GetWordSoFar(n)}
		} else {
			return nil
		}
	}

	if !n.hasChildren() {
		// we have found a terminal node, leaf
		return nil
	}

	var matchingChildren map[rune]*node
	ch := []rune(pattern)[0]
	if _, ok := n.children[ch]; !ok && ch != '*' && ch != '?' {
		return nil
	}
	if ch == '*' || ch == '?' {
		// Find all words under all children
		matchingChildren = n.children
	} else {
		// Find all words under a single matching child
		matchingChildren = map[rune]*node{ch: n.children[ch]}
	}

	var matches []string
	for _, child := range matchingChildren {
		if ch == '*' {
			for _, match := range child.FindStarMatching(pattern) {
				matches = append(matches, match)
			}
		}
		for _, match := range child.FindStarMatching(pattern[1:]) {
			matches = append(matches, match)
		}
	}
	return matches
}

// *y

//
//func FindAllMatchingStack(n *node, pattern string) []string {
//	var results []string
//	type item struct {
//		n          *node
//		patternIdx int
//	}
//	s := stack.New(item{n, 0})
//	for !s.Empty() {
//		i, _ := s.Pop()
//		if i.patternIdx == len(pattern) {
//			if i.n.isEnd {
//				results = append(results, GetWordSoFar(n))
//			}
//			continue
//		}
//		ch := pattern[i.patternIdx]
//	}
//	var matches []string
//	if pattern == "" && n.isEnd {
//		return []string{GetWordSoFar(n)}
//	}
//	if n.isEnd {
//		return nil
//	}
//
//	return matches
//}

//var first rune
//for _,c := range str {
//first = c
//break
//}

//
//func main() {
//	s := "世界 Hello"
//	r, size := utf8.DecodeRuneInString(s)
//
//	fmt.Printf("First rune: %c (Unicode point: %#U, size in bytes: %d)\n", r, r, size)
//
//	// Example with ASCII string
//	s2 := "Hello"
//	r2, size2 := utf8.DecodeRuneInString(s2)
//	fmt.Printf("First rune: %c (Unicode point: %#U, size in bytes: %d)\n", r2, r2, size2)
//}

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

func (t *Trie) FindLongestOneCharAtATime() string {
	if t == nil || t.Root == nil || t.Root.children == nil {
		return ""
	}
	output := ""
	s := stack.New(t.Root)
	for !s.Empty() {
		i, _ := s.Pop()
		if i.isEnd {
			for _, child := range i.children {
				s.Push(child)
			}
		}
	}
	return output
}
