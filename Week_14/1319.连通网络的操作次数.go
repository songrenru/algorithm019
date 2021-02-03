/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
 */

// @lc code=start
func makeConnected(n int, connections [][]int) int {
	lines := len(connections)
	if lines < n-1 {
		return -1
	}

	u := newUF(n)
	for _, connection := range connections {
		u.Union(connection[0], connection[1])
	}
	return u.Count() - 1
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
