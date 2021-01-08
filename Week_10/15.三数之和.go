/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)
	lastIdx := n - 3
	res := [][]int{}
	for i := 0; i <= lastIdx; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			temp := nums[i] + nums[l] + nums[r]
			if temp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l + 1] {
					l++
				}
				for l < r && nums[r] == nums[r - 1] {
					r--
				}
				l++
				r--
			} else if temp > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
// @lc code=end

