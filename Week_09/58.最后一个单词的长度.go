/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
    l := len(s)
    // 先排除掉最后的空str
    for i := l - 1; i >= 0; i-- {
        if s[i] == ' ' {
            l--
        } else {
            break
        }
    }
    // 再次从后面开始看
    count := 0
    for i := l - 1; i >= 0; i-- {
        if s[i] == ' ' {
            break
        }
        count++
    }
    return count
}
// @lc code=end

