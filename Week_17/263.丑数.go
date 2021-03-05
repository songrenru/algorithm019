/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	s := [3]int{2, 3, 5}
	for {
		if num == 1 {
			return true
		}
		next := false
		for _, v := range s {
			if num%v == 0 {
				num /= v
				next = true
				break
			}
		}
		if !next {
			return false
		}
	}
}

// @lc code=end

