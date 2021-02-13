/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	// dfs
	// union find
	rows, cols := len(grid), len(grid[0])
	count1 := 0
	total := rows*cols
	uf := newUF(total)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				count1++
				if i - 1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(i*rows+j, (i-1)*rows+j)
				}
				if i + 1 < rows && grid[i+1][j] == '1' {
					uf.Union(i*rows+j, (i+1)*rows+j)
				}
				if j - 1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*rows+j, i*rows+j-1)
				}
				if j + 1 < cols && grid[i][j+1] == '1' {
					uf.Union(i*rows+j, i*rows+j+1)
				}
			}
		}
	}
	return uf.Count() - (total - count1)
}


type UF struct {
	union []int
	count int
}

func newUF(cap int) *UF {
	uf := UF{
		union: make([]int, cap),
		count: cap,
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
	u.count--
}

func (u *UF) Connected(x, y int) bool {
	rootX := u.find(x)
	rootY := u.find(y)
	return rootX == rootY
}

func (u *UF) Count() int {
	return u.count
}

func (u *UF) find(x int) int {
	root := x
	for root != u.union[root] {
		root = u.union[root]
	}
	// 路径压缩
	for x != root {
		tmp := u.union[x]
		u.union[x] = root
		x = tmp
	}
	return root
}
// @lc code=end

