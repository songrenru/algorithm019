/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	// // recursion
	// if root == nil {
	// 	return nil
	// }
	// res := []int{}
	// res = append(res, root.Val)
	// res = append(res, preorderTraversal(root.Left)...)
	// res = append(res, preorderTraversal(root.Right)...)
	// return res

	// iteration
	res := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
// @lc code=end

