/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	// // recursion terminator
	// if root == nil {
	// 	return true
	// }
	// // process current logic
	// if root.Left != nil && root.Left.Val >= root.Val {
	// 	return false
	// }
	// if root.Right != nil && root.Right.Val <= root.Val {
	// 	return false
	// }
	// // dirll down
	// if !isValidBST(root.Left) {
	// 	return false
	// }
	// if !isValidBST(root.Right) {
	// 	return false
	// }
	// // reverse status
	// return true
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
// @lc code=end

