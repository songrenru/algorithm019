/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 暴力，按宽度枚举
	// 暴力，按高度枚举
	// 按高度枚举,单调栈 + sentinel

	// sentinel
	newHeights := []int{0}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)
	
	l := len(newHeights)
	stack := []int{0}
	size := 1
	area := 0
	for i := 1; i < l; i++ {
		for newHeights[i] < newHeights[stack[size - 1]] {
			h := newHeights[stack[size - 1]]
			l := stack[size - 2]
			curArea := h * (i - l - 1)
			if curArea > area {
				area = curArea
			}
			size--
			stack = stack[:size]
		}
		stack = append(stack, i)
		size++
	}
	return area
}
// @lc code=end

