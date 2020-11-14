/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	childrens := []*Node{root}
	
	for len(childrens) != 0 {
		idx := len(res)
		res = append(res, []int{})
		newChildrens := []*Node{}
		for _, node := range childrens {
			res[idx] = append(res[idx], node.Val)
			newChildrens = append(newChildrens, node.Children...)
		}
		childrens = newChildrens
	}

	return res
}
// @lc code=end

