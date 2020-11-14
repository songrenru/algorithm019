/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	idxs := make(map[string]int) // 用来存储索引

	for _, str := range strs {
		temp := sortByte(str)
		sort.Sort(temp)
		res := string(temp)
		if idx, ok := idxs[res]; !ok {
			idxs[res] = len(ans)
			ans = append(ans, []string{str})
		} else {
			ans[idx] = append(ans[idx], str)
		}
	}
	return ans
}

type sortByte []byte

func (a sortByte) Len() int           { return len(a) }
func (a sortByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByte) Less(i, j int) bool { return a[i] < a[j] }
// @lc code=end

