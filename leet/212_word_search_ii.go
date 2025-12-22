package leet

type TrieNode struct {
	next [26]*TrieNode
	word string
}

func insert(root *TrieNode, w string) {
	cur := root
	for i := 0; i < len(w); i++ {
		idx := w[i] - 'a'
		if cur.next[idx] == nil {
			cur.next[idx] = &TrieNode{}
		}
		cur = cur.next[idx]
	}
	cur.word = w
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	root := &TrieNode{}
	for _, w := range words {
		insert(root, w)
	}

	rows, cols := len(board), len(board[0])
	res := make([]string, 0)

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		ch := board[r][c]
		if ch == '#' {
			return
		}

		child := node.next[ch-'a']
		if child == nil {
			return
		}

		if child.word != "" {
			res = append(res, child.word)
			child.word = ""
		}

		board[r][c] = '#'
		if r > 0 {
			dfs(r-1, c, child)
		}
		if r+1 < rows {
			dfs(r+1, c, child)
		}
		if c > 0 {
			dfs(r, c-1, child)
		}
		if c+1 < cols {
			dfs(r, c+1, child)
		}
		board[r][c] = ch
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return res
}
