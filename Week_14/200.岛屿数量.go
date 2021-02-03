/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	u := newUF(rows * cols)
	count1 := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count1++
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					u.Union(i*cols+j, (i-1)*cols+j)
				}
				if i+1 < rows && grid[i+1][j] == '1' {
					u.Union(i*cols+j, (i+1)*cols+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					u.Union(i*cols+j, i*cols+j-1)
				}
				if j+1 < cols && grid[i][j+1] == '1' {
					u.Union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}
	return u.Count() - (rows*cols - count1)
}

type UF struct {
	union []int
	count int
}

func newUF(cap int) *UF {
	uf := UF{
		make([]int, cap),
		cap,
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

func (u *UF) Count() int {
	return u.count
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
