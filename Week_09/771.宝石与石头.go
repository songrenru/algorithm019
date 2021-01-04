/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
    jMap := map[byte]bool{}
    jL := len(jewels)
    for i := 0; i < jL; i++ {
        jMap[jewels[i]] = true
    }
    count := 0
    l := len(stones)
    for i := 0; i < l; i++ {
        if jMap[stones[i]] {
            count++
        }
    }
    return count
}
// @lc code=end

