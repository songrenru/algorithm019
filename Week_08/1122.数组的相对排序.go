/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	res := []int{}
	l1 := len(arr1)
	l2 := len(arr2)
	for i := 0; i < l2; i++ {
		idx := sort.SearchInts(arr1, arr2[i])
		for idx < l1 && arr1[idx] == arr2[i] {
			res = append(res, arr2[i])
			idx++
		}
	}
	
	sort.Ints(arr2)
	for i := 0; i < l1; i++ {
		idx := sort.SearchInts(arr2, arr1[i])
		if idx == l2 || arr1[i] != arr2[idx] {
			res = append(res, arr1[i])
		}
	}
	return res
}
// @lc code=end

