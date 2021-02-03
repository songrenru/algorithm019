/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	u := newUF(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}
	return u.Count()
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
