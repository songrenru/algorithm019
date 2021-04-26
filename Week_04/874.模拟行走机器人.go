/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	directs := [4][2]int{
		{0, 1},  // noth = iota
		{1, 0},  // east
		{0, -1}, // south
		{-1, 0}, // west
	}
	obstaclesSet := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstaclesSet[[2]int{obstacle[0], obstacle[1]}] = true
	}

	res := 0
	direct := 0 // noth = 0
	x, y := 0, 0
	for _, cmd := range commands {
		if cmd == -1 {
			direct = (direct + 1) % 4
			continue
		} else if cmd == -2 {
			direct = (direct + 3) % 4
			continue
		} else {
			xChg, yChg := directs[direct][0], directs[direct][1]
			for i := 0; i < cmd; i++ {
				x += xChg
				y += yChg
				if obstaclesSet[[2]int{x, y}] {
					x -= xChg
					y -= yChg
					break
				}
			}
		}
		distance := x*x + y*y
		if res < distance {
			res = distance
		}
	}

	return res
}

// @lc code=end

