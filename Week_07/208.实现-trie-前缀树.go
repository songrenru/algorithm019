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
	lens := len(word)
	node := this
	for i := 0; i < lens; i++ {
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
	lens := len(word)
	node := this
	for i := 0; i < lens; i++ {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	lens := len(prefix)
	node := this
	for i := 0; i < lens; i++ {
		idx := prefix[i] - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

