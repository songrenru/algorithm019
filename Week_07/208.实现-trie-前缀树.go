/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	next [26]*Trie
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	l := len(word)
	for i := 0; i < l; i++ {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			node.next[idx] = new(Trie)
		}
		node = node.next[idx]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	if node == nil {
		return false
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchPrefix(prefix)
	return node != nil
}


func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	l := len(prefix)
	for i := 0; i < l; i++ {
		idx := prefix[i] - 'a'
		if node.next[idx] == nil {
			return nil
		}
		node = node.next[idx]
	}
	return node
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

