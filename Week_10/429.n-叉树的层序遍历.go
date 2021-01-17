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
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		curRes := []int{}
		for i := 0; i < size; i++ {
			curRes = append(curRes, q[i].Val)
			for _, v := range q[i].Children {
				if v != nil {
					q = append(q, v)
				}
			}
		}
		res = append(res, curRes)
		q = q[size:]
	}
	return res
}
// @lc code=end

