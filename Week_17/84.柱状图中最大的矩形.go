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

	stack := []int{0}
	size := 1
	l := len(newHeights)
	maxArea := 0
	for i := 1; i < l; i++ {
		// 打掉前面比当前大的元素（单调栈）
		for newHeights[i] < newHeights[stack[size-1]] {
			height := newHeights[stack[size-1]]
			leftIdx := stack[size-2] // 左边第一条肯定《=当前height
			width := i - leftIdx - 1
			area := height * width
			if area > maxArea {
				maxArea = area
			}
			size--
			stack = stack[:size]
		}
		size++
		stack = append(stack, i)
	}

	return maxArea
}

// @lc code=end

