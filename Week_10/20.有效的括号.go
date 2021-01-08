/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	l := len(s)
	if l & 1 == 1 {
		return false
	}

	maps := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	size := 0
	for i := 0; i < l; i++ {
		if c, ok := maps[s[i]]; ok {
			if size == 0 || stack[size] != c {
				return false
			}
			size--
			stack = stack[:size]
		} else {
			stack = append(stack, s[i])
			size++
		}
	}
	return size == 0
}
// @lc code=end

