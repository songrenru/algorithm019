/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	// bfs
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		max := math.MinInt64
		newNodes := []*TreeNode{}
		for i := 0; i < size; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				newNodes = append(newNodes, queue[i].Left)
			}
			if queue[i].Right != nil {
				newNodes = append(newNodes, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = newNodes
	}

	return res
}
// @lc code=end

