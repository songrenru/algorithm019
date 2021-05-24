/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	dp := treeDp(root)
	return max(dp[0], dp[1])
}

func treeDp(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{}
	}

	left := treeDp(root.Left)
	right := treeDp(root.Right)
	res := [2]int{}
	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = root.Val + left[0] + right[0]
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

