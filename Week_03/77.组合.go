/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	if k <= 0 || k > n {
		return nil
	}
	res := [][]int{}
	temp := []int{}
	
	var helper func(i int)
	helper = func(i int)  {
		// recursion terminator
		if len(temp) + (n - i + 1) < k {
			return
		}
		// process current logic
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			res = append(res, comb)
			return
		}
		// drill down
		temp = append(temp, i)
		helper(i + 1)
		temp = temp[:len(temp)-1]
		// reverse state
		helper(i + 1)
	}
	helper(1)
	return res
}
// @lc code=end

