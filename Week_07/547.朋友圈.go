/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
func findCircleNum(M [][]int) int {
	n := len(M)
	uf := newUF(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count()
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
	for u.union[root] != root {
		root = u.union[root]
	}
	for u.union[x] != root { // 路径压缩
		temp := u.union[x]
		u.union[x] = root
		x = temp
	}
	return root
}

// @lc code=end

