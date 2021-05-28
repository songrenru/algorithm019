/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	l := len(word)
	root := this
	for i := 0; i < l; i++ {
		if root.next[word[i]-'a'] == nil {
			newRoot := Constructor()
			root.next[word[i]-'a'] = &newRoot
		}
		root = root.next[word[i]-'a']
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	l := len(word)
	root := this
	for i := 0; i < l; i++ {
		if root.next[word[i]-'a'] == nil {
			return false
		}
		root = root.next[word[i]-'a']
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	l := len(prefix)
	root := this
	for i := 0; i < l; i++ {
		if root.next[prefix[i]-'a'] == nil {
			return false
		}
		root = root.next[prefix[i]-'a']
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

