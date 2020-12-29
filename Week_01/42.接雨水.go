/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// 暴力，按“高度”枚举，三种情况
	// 按“高度”枚举 + dp构造左右边界
	// 按“高度”枚举 + dp思想，双指针替代
	l := len(height)
	if l <= 2 {
		return 0
	}
	
	// dp构造左右边界
	maxLeft := 0
	maxRight := 0
	left := 1
	right := l - 2
	
	vol := 0
	for i := 1; i < l - 1; i++ {
		if height[left - 1] < height[right + 1] { // 从左到右更
			maxLeft = max(maxLeft, height[left - 1])
			min := maxLeft
			if min > height[left] {
				vol += min - height[left]
			}
			left++
		} else { // 从右到左更
			maxRight = max(maxRight, height[right + 1])
			min := maxRight
			if min > height[right] {
				vol += min - height[right]
			}
			right--
		}
	}
	return vol
}

func max(i, j int) int {
	if i > j { return i }
	return j
}
// @lc code=end

