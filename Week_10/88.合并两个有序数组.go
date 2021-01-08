/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	t := len(nums1) - 1
	for m > 0 && n > 0 {
		if nums1[m - 1] > nums2[n - 1] {
			m--
			nums1[t] = nums1[m]
		} else {
			n--
			nums1[t] = nums2[n]
		}
		t--
	}
	for i := 0; i < n; i++ {
		nums1[i] = nums2[i]
	}
}
// @lc code=end

