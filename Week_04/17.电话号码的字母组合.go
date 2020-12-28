/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	lens := len(digits)
	if lens == 0 {
		return nil
	}

	maps := make(map[byte][]byte)
	maps['2'] = []byte{'a', 'b', 'c'}
	maps['3'] = []byte{'d', 'e', 'f'}
	maps['4'] = []byte{'g', 'h', 'i'}
	maps['5'] = []byte{'j', 'k', 'l'}
	maps['6'] = []byte{'m', 'n', 'o'}
	maps['7'] = []byte{'p', 'q', 'r', 's'}
	maps['8'] = []byte{'t', 'u', 'v'}
	maps['9'] = []byte{'w', 'x', 'y', 'z'}

	res := []string{}
	var dfs func(i int, path []byte)
	dfs = func(i int, path []byte) {
		if i == lens {
			res = append(res, string(path))
			return
		}
	
		for _, c := range maps[digits[i]] {
			path = append(path, c)
			dfs(i + 1, path)
			path = path[:len(path) - 1]
		}
	}
	dfs(0, []byte{})
	return res
}


// @lc code=end

