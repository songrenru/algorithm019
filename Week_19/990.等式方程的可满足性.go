/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	uf := [26]byte{}
	for i := byte(1); i < 26; i++ {
		uf[i] = i
	}

	// 先处理相等的
	for _, equation := range equations {
		if equation[1] == '=' {
			union(&uf, equation[0]-'a', equation[3]-'a')
		}
	}

	// 判定!=的情况是否矛盾
	for _, equation := range equations {
		if equation[1] == '!' {
			if find(&uf, equation[0]-'a') == find(&uf, equation[3]-'a') {
				return false
			}
		}
	}
	return true
}

func union(uf *[26]byte, x, y byte) {
	rootX := find(uf, x)
	rootY := find(uf, y)
	if rootX == rootY {
		return
	}
	uf[rootX] = rootY
}

func find(uf *[26]byte, x byte) byte {
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

