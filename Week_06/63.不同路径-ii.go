/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r := len(obstacleGrid)
	c := len(obstacleGrid[0])
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else if i == r - 1 && j == c - 1 {
				obstacleGrid[i][j] = 1 // 目标点作为起始点，特殊处理，也为后续步骤做铺垫
			} else {
				// 判定右边、下边有没有障碍物，有则该路径=0
				down, right := 0, 0
				if !(j == c - 1 || obstacleGrid[i][j+1] == 0) {
					right = obstacleGrid[i][j+1]
				}
				if !(i == r - 1 || obstacleGrid[i+1][j] == 0) {
					down = obstacleGrid[i+1][j]
				}
				obstacleGrid[i][j] = down + right
			}
		}
	}
	return obstacleGrid[0][0]
}
// @lc code=end

