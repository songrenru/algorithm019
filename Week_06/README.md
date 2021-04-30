# 学习笔记
## 动态规划
1. dynamic programming动态递推
2. 定义:“Simplifying a complicated problem by breaking it down into 
simpler sub-problems” 
(in a recursive manner)
3. Divide & Conquer + Optimal substructure 
 分治 + 最优子结构
## 关键点：
1. dp 和 递归或者分治 没有根本上的区别 （关键看有无最优子结构）
2. 共性：找到重复子问题
3. 最优子结构、中途可以淘汰次优解
## 解题三步骤：
1. 最优子结构 opt[n] = best_of(opt[n-1], opt[n-2], …)
2. 储存中间状态：opt[i]
3. 递推公式（美其名曰：状态转移方程或者 DP 方程）
   1. Fib: opt[i] = opt[n-1] + opt[n-2]
   2. 二维路径：opt[i][j] = opt[i+1][j] + opt[i][j+1]
