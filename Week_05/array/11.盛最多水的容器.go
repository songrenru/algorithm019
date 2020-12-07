/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 1. 双循环暴力求解
	// 2. 前后双指针（宽在缩小，高度替换最小的边才可能容积变大）
	cap := 0
	i, j := 0, len(height) - 1
	for i < j {
		temp := (j - i) * min(height[i], height[j])
		if temp > cap {
			cap = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return cap
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
// @lc code=end

