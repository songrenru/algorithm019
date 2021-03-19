/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	// 递归
	// 广度优先遍历
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}
		}
		queue = queue[size:]
	}
	return depth
}

// @lc code=end

