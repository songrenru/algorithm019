/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i, val := range nums {
        idx, ok := dict[target - val]
        if ok {
            return []int{i, idx}
        }
        dict[val] = i
    }
    return nil
}
// m1:暴力破解
// m2:两遍循环
// m3:一遍循环
// @lc code=end

