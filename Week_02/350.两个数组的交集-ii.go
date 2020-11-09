/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	map1 := make(map[int]int)
	for _, v := range nums1 {
		map1[v]++
	}
	intersection := []int{}
	for _, v := range nums2 {
		if map1[v] > 0 {
			intersection = append(intersection, v)
			map1[v]--
		}
	}
	return intersection
}
// @lc code=end

