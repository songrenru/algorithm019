/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    res := []int{}
	var helper func(node *Node)
	helper = func (node *Node) {
		if node == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			helper(node.Children[i])
		}
		res = append(res, node.Val)
	}
	helper(root)
	return res
}
// @lc code=end

