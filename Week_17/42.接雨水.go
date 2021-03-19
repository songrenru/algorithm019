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
	dpLeft := make([]int, l)
	dpLeft[0] = 0
	dpRight := make([]int, l)
	dpRight[l-1] = 0
	for i := 1; i < l; i++ {
		dpLeft[i] = max(dpLeft[i-1], height[i-1])
	}
	for i := l - 2; i >= 0; i-- {
		dpRight[i] = max(dpRight[i+1], height[i+1])
	}

	vol := 0
	for i := 1; i < l-1; i++ {
		minBound := min(dpLeft[i], dpRight[i])
		if height[i] < minBound {
			vol += minBound - height[i]
		}
	}
	return vol
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

