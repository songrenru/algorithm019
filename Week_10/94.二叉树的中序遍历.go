/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	// 递归
	// 迭代
	stack := []*TreeNode{}
	res := []int{}
	// root != nil，即兼容初始root == nil的情况
	// 又实现的树的回溯
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, root.Val)
		root = root.Right // 配合第一层for的root != nil,实现回溯
	}
	return res
}
// @lc code=end

