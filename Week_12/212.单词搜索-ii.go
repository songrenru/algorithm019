/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	// 构造字典树
	trie := new(Trie)
	for _, word := range words {
		trie.Insert(word)
	}
	// 定义方向
	directs := [4][2]int{
		{-1, 0}, // up
		{1, 0}, // down 
		{0, -1}, // left
		{0, 1}, // right
	}
	rows, cols := len(board), len(board[0])
	res := []string{}

	var dfs func(i, j int, node *Trie) 
	dfs = func(i, j int, node *Trie) {
		// 越界检测
		if i < 0 || j < 0 || i == rows || j == cols {
			return
		}
		// 不能重复
		c := board[i][j]
		if c == '#' {
			return
		}
		// 单词前缀检测
		newNode = node.next[c - 'a']
		if newNode == nil {
			return
		}
		if newNode.isEnd {
			// 剪枝
			newNode.isEnd = false
			res = append(res, newNode.word)
		}
		// 继续遍历
		board[i][j] = '#'
		for k := 0; k < 4; k++ {
			dfs(i + directs[k][0], j + directs[k][1], newNode)
		}
		// 恢复
		board[i][j] = c
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, trie)
		}
	}
	return res
}

type Trie struct {
	next [26]*Trie
	isEnd bool
	word string
}

func (this *Trie) Insert(word string)  {
	node := this
	l := len(word)
	for i := 0; i < l; i++ {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			node.next[idx] = new(Trie)
		}
		node = node.next[idx]
	}
	node.isEnd = true
	node.word = word
}
// @lc code=end

