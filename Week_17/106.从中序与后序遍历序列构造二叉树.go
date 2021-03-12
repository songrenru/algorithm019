/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// recursion terminator
	if len(inorder) == 0 {
		return nil
	}

	// process current logic
	lastIdx := len(postorder) - 1
	val := postorder[lastIdx]
	node := &TreeNode{Val: val}

	if lastIdx == 0 {
		return node
	}

	mid := findMid(inorder, val)

	node.Left = buildTree(inorder[0:mid], postorder[0:mid])
	node.Right = buildTree(inorder[mid+1:], postorder[mid:lastIdx])

	// revert status if needed
	// return
	return node
}

func findMid(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}

// @lc code=end

