/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	directs := [4][2]int{
		{-1, 0}, // top
		{1, 0}, // down
		{0, -1}, // left
		{0, 1}, // right
	}
	row, col := len(board), len(board[0])
	set := map[string]bool{}

	var dfs func(i, j int, path []byte, node *Trie)
	dfs = func(i, j int, path []byte, node *Trie)  {
		// 越界检测
		if i < 0 || j < 0 || i == row || j == col {
			return
		}
	
		// 单词检测
		c := board[i][j]
		if c == '#' || node.next[c - 'a'] == nil {
			return
		}
	
		newNode := node.next[c - 'a']
		path = append(path, c)
		if newNode.isEnd {
			set[string(path)] = true
		}
	
		// dfs + backtrace
		board[i][j] = '#'
		for k := 0; k < 4; k++ {
			dfs(i + directs[k][0], j + directs[k][1], path, newNode)
		}
		board[i][j] = c
		path = path[:len(path) - 1]
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dfs(i, j, []byte{}, &trie)
		}
	}

	res := []string{}
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}



type Trie struct {
	next [26]*Trie
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	lens := len(word)
	node := this
	for i := 0; i < lens; i++ {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			node.next[idx] = new(Trie)
		}
		node = node.next[idx]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	lens := len(word)
	node := this
	for i := 0; i < lens; i++ {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	lens := len(prefix)
	node := this
	for i := 0; i < lens; i++ {
		idx := prefix[i] - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return true
}
// @lc code=end

