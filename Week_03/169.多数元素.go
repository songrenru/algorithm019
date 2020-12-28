/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	return bs(nums, 0, len(nums) - 1)
}

func bs(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := (hi - lo) / 2 + lo
	left := bs(nums, lo, mid)
	right := bs(nums, mid + 1, hi)

	leftCount := countInRange(nums, lo, mid, left)
	rightCount := countInRange(nums, mid + 1, hi, right)

	if leftCount > rightCount {
		return left
	}
	return right
}

func countInRange(nums []int, lo, hi, num int) int {
	count := 0
	for i := lo; i <= hi; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}
// @lc code=end

