/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	arr := byteSlice(s)
	sort.Sort(arr)
	i := 0
	l := len(s)
	res := []byte{}
	for i < l {
		cursor := ''
		for j := 0; j < l; j++ {
			if arr[j] == '0'{
				continue
			} else if cursor == '' || arr[j] != cursor {
				cursor = arr[j]
				res = append(res, arr[j])
				arr[j] = '0'
			}
		}

	}
}

type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }
// @lc code=end

