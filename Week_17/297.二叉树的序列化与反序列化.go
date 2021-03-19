/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	str := []string{}
	var rSerialize func(root *TreeNode)
	rSerialize = func(root *TreeNode) {
		if root == nil {
			str = append(str, "null")
			return
		}
		str = append(str, strconv.Itoa(root.Val))
		rSerialize(root.Left)
		rSerialize(root.Right)
	}
	rSerialize(root)
	return strings.Join(str, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		this.l = append(this.l, l[i])
	}
	return this.rDeserialize()
}

func (this *Codec) rDeserialize() *TreeNode {
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.rDeserialize()
	root.Right = this.rDeserialize()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

