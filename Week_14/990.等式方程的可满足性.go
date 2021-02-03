/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	// 计数
	u := newUF()
	// 先处理相等的
	for _, equation := range equations {
		if equation[0] != equation[3] && equation[1] == '=' {
			u.Union(equation[0]-'a', equation[3]-'a')
		}
	}
	// 判定!=的情况是否矛盾
	for _, equation := range equations {
		if equation[1] == '!' && u.Connected(equation[0]-'a', equation[3]-'a') {
			return false
		}
	}
	return true
}

type UF struct {
	union [26]byte
}

func newUF() *UF {
	uf := UF{}
	for i := byte(0); i < 26; i++ {
		uf.union[i] = i
	}
	return &uf
}

func (u *UF) Union(x, y byte) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	u.union[rootX] = rootY
	// u.count--
}

func (u *UF) Connected(x, y byte) bool {
	rootX := u.find(x)
	rootY := u.find(y)
	return rootX == rootY
}

func (u *UF) find(x byte) byte {
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
