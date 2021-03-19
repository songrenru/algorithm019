/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 暴力 + 去重，O(n^3)
	// 排序 + 二分查找 + 去重，O(N^2 * logN)
	// 排序 + 双指针剪枝去重，O(N^2)
	l := len(nums)
	if l < 3 {
		return nil
	}

	sort.Ints(nums)
	lastIdx := l - 2
	res := [][]int{}
	for i := 0; i < lastIdx; i++ {
		// 剪枝
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针查找
		search := -nums[i]
		left, right := i+1, lastIdx+1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < search {
				// for left < right && nums[left] == nums[left+1] {
				// 	left++
				// }
				left++
			} else if sum > search {
				// for left < right && nums[right] == nums[right-1] {
				// 	right--
				// }
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			}
		}

	}
	return res
}

// @lc code=end

