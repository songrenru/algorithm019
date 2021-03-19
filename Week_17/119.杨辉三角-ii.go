/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	// dp[i] = dp[i - 1] + dp[i]
	row := make([]int, rowIndex+1)
	for r := 0; r <= rowIndex; r++ {
		for c := r; c >= 0; c-- {
			if c == 0 || c == r {
				row[c] = 1
			} else {
				row[c] = row[c] + row[c-1]
			}
		}
	}
	return row
}

// @lc code=end

