/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// 回溯(for|线性)
	// 迭代(2^n)
	l := len(nums)
	t := 1 << l
	res := [][]int{}
	for mask := 0; mask < t; mask++ {
		set := []int{}
		for j, num := range nums {
			if mask>>j&1 == 1 {
				set = append(set, num)
			}
		}
		res = append(res, set)
	}
	return res
}

// @lc code=end

