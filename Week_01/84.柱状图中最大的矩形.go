/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 暴力，按“宽”枚举
	// 暴力，按“高”枚举
	// 单调栈 + 哨兵
	
	l := len(heights)
	// 哨兵
	newHeights := []int{0}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)
	l += 2

	stack := []int{0}

	maxArea := 0
	for i := 1; i < l; i++ {
		for newHeights[i] < newHeights[stack[len(stack) - 1]] {
			curHeight := newHeights[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1] // height出栈
			curWidth := i - stack[len(stack) - 1] - 1

			curArea := curHeight * curWidth
			maxArea = max(maxArea, curArea)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func max(i, j int) int {
	if i > j { return i }
	return j
}
// @lc code=end

