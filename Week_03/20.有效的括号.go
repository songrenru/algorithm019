/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	n := len(s)
    if n & 1 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	stack := []byte{}
	size := 0
    for i := 0; i < n; i++ {
        if pairs[s[i]] > 0 {
            if size == 0 || stack[size-1] != pairs[s[i]] {
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

