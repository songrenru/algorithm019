/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	// recursion terminator
	if root == nil {
		return nil
	}
	// process current logic
	root.Left, root.Right = root.Right, root.Left
	// dirll down
	invertTree(root.Left)
	invertTree(root.Right)
	// reverse status
	return root

}
// @lc code=end

