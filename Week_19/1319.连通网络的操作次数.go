/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
 */

// @lc code=start
func makeConnected(n int, connections [][]int) int {
	// n-1根线缆即可实现所有计算机互联
	// 移动count-1根线缆，即可连通整个网络
	lines := len(connections)
	if lines < n-1 {
		return -1
	}

	// 构建连通分量
	u := newUF(n)
	for _, connection := range connections {
		u.Union(connection[0], connection[1])
	}
	return u.count - 1
}

type uf struct {
	union []int
	count int
}

func newUF(cap int) uf {
	union := make([]int, cap)
	for i := 1; i < cap; i++ {
		union[i] = i
	}
	return uf{
		union: union,
		count: cap,
	}
}

func (u *uf) Union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	u.union[rootX] = rootY
	u.count--
}

func (u *uf) find(x int) int {
	root := x
	for root != u.union[root] {
		root = u.union[root]
	}
	// 路径压缩
	for root != x {
		tmp := u.union[x]
		u.union[x] = root
		x = tmp
	}
	return root
}

// @lc code=end

