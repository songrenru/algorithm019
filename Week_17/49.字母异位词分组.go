/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	idxMap := map[string]int{}
	for _, str := range strs {
		tmp := sortByte(str)
		sort.Sort(tmp)
		key := string(tmp)
		if idx, ok := idxMap[key]; !ok {
			idxMap[key] = len(res)
			res = append(res, []string{str})
		} else {
			res[idx] = append(res[idx], str)
		}
	}
	return res
}

type sortByte []byte

func (a sortByte) Len() int           { return len(a) }
func (a sortByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByte) Less(i, j int) bool { return a[i] < a[j] }

// @lc code=end

