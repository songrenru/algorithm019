/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Ints(nums)
	lastIdx := n - 4
	res := [][]int{}
	for j := 0; j <= lastIdx; j++ {
		if j > 0 && nums[j] == nums[j - 1] {
			continue
		}
		curTar := target - nums[j]

		for i := j + 1; i < n; i++ {
			// if nums[i] > curTar {
			// 	break
			// }
			if i > j + 1 && nums[i] == nums[i - 1] {
				continue
			}
			
			l := i + 1
			r := n - 1
			for l < r {
				temp := nums[i] + nums[l] + nums[r]
				if temp == curTar {
					res = append(res, []int{nums[j], nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l + 1] {
						l++
					}
					for l < r && nums[r] == nums[r - 1] {
						r--
					}
					l++
					r--
				} else if temp > curTar {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
// @lc code=end

