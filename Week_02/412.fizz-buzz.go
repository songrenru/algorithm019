/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
    res := make([]string, n)
    for i := 1; i <= n; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
        if div3 {
            if div5 {
                res[i-1] = "FizzBuzz"
            } else {
                res[i-1] = "Fizz"
            }
        } else if div5 {
			res[i-1] = "Buzz"
        } else {
			res[i-1] = strconv.Itoa(i)
        }
    }
    return res
}
// @lc code=end

