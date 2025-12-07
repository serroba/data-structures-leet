package trie

type node struct {
	isEnd    bool
	children map[rune]*node
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

func (t *Trie) Insert(word string) {
	current := t.Root
	for _, ch := range word {
		if _, ok := current.children[ch]; !ok {
			current.children[ch] = &node{children: make(map[rune]*node)}
		}
		current = current.children[ch]
	}
	current.isEnd = true
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
	if !t.SearchWord(word) {
		return
	}
	current := t.Root

	for _, ch := range word {
		if _, ok := current.children[ch]; ok {
			current = current.children[ch]
			if current.isEnd && current.children != nil {
				current.isEnd = false
			}
		}
	}
}
