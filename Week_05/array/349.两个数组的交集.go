/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	set1 := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	
	res := []int{}
	l1 := len(set1)
	l2 := len(nums2)
	for i := 0; i < l2 && l1 > 0; i++ {
		if set1[nums2[i]] {
			res = append(res, nums2[i])
			delete(set1, nums2[i])
			l1--
		}
	}
	return res

}
// @lc code=end

