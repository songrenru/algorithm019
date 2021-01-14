/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	l := len(height)
	if l <= 2 {
		return 0
	}

	// dp 构造左右边界
	dp_l := make([]int, l)
	dp_l[0] = height[0]
	for i := 1; i < l; i++ {
		dp_l[i] = max(dp_l[i - 1], height[i])
	}

	dp_r := make([]int, l)	
	dp_r[l - 1] = height[l - 1]
	for i := l - 2; i >= 0; i-- {
		dp_r[i] = max(dp_r[i + 1], height[i])
	}

	vol := 0
	for i := 1; i < l - 1; i++ {
		curHeight := min(dp_l[i], dp_r[i])
		if curHeight > height[i] {
			vol += curHeight - height[i]
		}
	}
	return vol
}

func max(i, j int) int {
	if i > j { return i }
	return j
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

