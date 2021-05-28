/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	// union find + dfs
	// 技巧： 构造一个锚点，和边界上的'O'联通
	rows, cols := len(board), len(board[0])
	if rows <= 2 || cols <= 2 {
		return
	}

	// 构造锚点
	anchor := rows * cols
	uf := make([]int, anchor+1)
	for i := 1; i <= anchor; i++ {
		uf[i] = i
	}

	// 构建union find
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				// 边界上，关联锚点
				if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
					union(&uf, genIdx(i, j, cols), anchor)
				}
				// 联通右边、下边的位子
				if j+1 < cols && board[i][j+1] == 'O' {
					union(&uf, genIdx(i, j, cols), genIdx(i, j+1, cols))
				}
				if i+1 < rows && board[i+1][j] == 'O' {
					union(&uf, genIdx(i, j, cols), genIdx(i+1, j, cols))
				}
			}
		}
	}

	// 填充
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && !isConnected(&uf, genIdx(i, j, cols), anchor) {
				board[i][j] = 'X'
			}
		}
	}

}

func genIdx(i, j, cols int) int {
	return i*cols + j
}

func isConnected(uf *[]int, x, y int) bool {
	rootX, rootY := find(uf, x), find(uf, y)
	return rootX == rootY
}

func union(uf *[]int, x, y int) {
	rootX, rootY := find(uf, x), find(uf, y)
	if rootX == rootY {
		return
	}
	(*uf)[rootX] = rootY
}

func find(uf *[]int, x int) int {
	root := x
	for root != (*uf)[root] {
		root = (*uf)[root]
	}
	// 路径压缩
	for x != (*uf)[x] {
		tmp := (*uf)[x]
		(*uf)[x] = root
		x = tmp
	}
	return root
}

// @lc code=end

