/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	// bfs
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	

	for len(queue) != 0 {
		size := len(queue)
		eles := []int{}
		newNodes := []*TreeNode{}
		for i := 0; i < size; i++ {
			eles = append(eles, queue[i].Val)
			if queue[i].Left != nil {
				newNodes = append(newNodes, queue[i].Left)
			}
			if queue[i].Right != nil {
				newNodes = append(newNodes, queue[i].Right)
			}
		}
		res = append(res, eles)
		queue = newNodes
	}

	return res

	// dfs
}
// @lc code=end

