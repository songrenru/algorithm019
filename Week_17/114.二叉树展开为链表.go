/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	// 构建一个新的链表，然后再把root指过去
	if root == nil {
		return
	}
	newRoot := new(TreeNode)
	pre := newRoot
	var inorderTravel func(root *TreeNode)
	inorderTravel = func(root *TreeNode) {
		if root == nil {
			return
		}

		left, right := root.Left, root.Right

		// 清楚left、right，构建单链表形式
		root.Left, root.Right = nil, nil
		pre.Right = root
		pre = pre.Right

		inorderTravel(left)
		inorderTravel(right)
	}
	inorderTravel(root)

	// 把root指过去
	root.Right = newRoot.Right.Right
}

// @lc code=end

