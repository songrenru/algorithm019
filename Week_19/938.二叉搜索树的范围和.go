/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	var num int
	var preorderTraversal func(root *TreeNode)
	preorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		// <= low,右子树
		if root.Val < low {
			preorderTraversal(root.Right)
			return
		}
		// >= high，左子树
		if root.Val > high {
			preorderTraversal(root.Left)
			return
		}
		preorderTraversal(root.Left)
		num += root.Val
		preorderTraversal(root.Right)
	}
	preorderTraversal(root)
	return num
}

// @lc code=end

