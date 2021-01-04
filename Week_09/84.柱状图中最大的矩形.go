/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 单调栈 + 哨兵

	// sentinel
	l := len(heights)
	if l == 0 {
		return 0
	}
	// newHeights := []int{0}
	// newHeights = append(newHeights, heights)
	// newHeights = append(newHeights, 0)

	stack := []int{-1, 0} // 0是索引0
	for i := 1; i < l; i++ {
		if heights[i] < heights[len(stack) - 1] {
			size = len(stack) - 1
			curHeight := heights[stack[size]]
			stack = stack[:size]
			curWidth := i - 
		}
	}
}
// @lc code=end

