# 学习笔记
## 递归
1. 模板：
    1. recursion terminator
    2. process current logic
    3. drill down
    4. reverse the current level status if needed
2. 思维要点：
    1. 不要人肉进行递归（最大误区）
    2. 找到最近最简方法，将其拆分成可重复解决的问题（重复子问题）
    3. 数学归纳法的思维
## 分治、回溯（本质是递归，找重复性）
1. 分治divide conquer模板：
    1. recursion terminator
    2. prepare data
    3. conquer subproblems
    4. process and generate the final result
    5. reverse the current level status if needed
2. 回溯backtracking（又叫试探法）:采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案;而满足回溯条件的某个状态的点称为“回溯点”;回溯法通常用最简单的递归方法来实现。




