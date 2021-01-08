/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0
    i, j := 0, len(height)-1
    for i < j {
		if cap := min(height[i], height[j]) * (j - i); cap > max {
			max = cap
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
    }
    return max
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

