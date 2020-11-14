/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}
// @lc code=end

