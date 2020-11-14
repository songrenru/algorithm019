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
    if root == nil {
		return nil
	}
	res := []int{}
	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}
	res = append(res, root.Val)
	return res
}
// @lc code=end

