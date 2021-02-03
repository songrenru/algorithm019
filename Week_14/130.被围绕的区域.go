/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	rows := len(board)
	// 少于两行的，pass
	if rows <= 2 {
		return
	}
	cols := len(board[0])
	// 并查集锚点，将所有边界为O的，归属到一个集合
	anchor := rows * cols
	u := newUF(anchor + 1)

	// 构建联通区域
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				curNum := genNum(i, j, cols)
				if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
					u.Union(anchor, curNum)
				}
				// 每次只判定左边、上边元素，就可覆盖全部
				// 注意，由于只判定左、上，所以边界也要进行判定，不然会漏掉边界
				if i-1 >= 0 && board[i-1][j] == 'O' {
					u.Union(curNum, genNum(i-1, j, cols))
				}
				if j-1 >= 0 && board[i][j-1] == 'O' {
					u.Union(curNum, genNum(i, j-1, cols))
				}
			}
		}
	}

	// 判定联通区域
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && !u.Connected(anchor, genNum(i, j, cols)) {
				board[i][j] = 'X'
			}
		}
	}
}

func genNum(i, j, cols int) int {
	return i*cols + j
}

type UF struct {
	union []int
}

func newUF(cap int) *UF {
	uf := UF{
		make([]int, cap),
	}
	for i := 0; i < cap; i++ {
		uf.union[i] = i
	}
	return &uf
}

func (u *UF) Union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	u.union[rootX] = rootY
}

func (u *UF) Connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UF) find(x int) int {
	root := x
	for u.union[root] != root {
		root = u.union[root]
	}
	for x != root {
		tmp := u.union[x]
		u.union[x] = root
		x = tmp
	}
	return root
}

// @lc code=end
