# 学习笔记

## 精通一个领域（3）
### chunk it up 
1. 构造知识树、网络（脑图）
    1. 数据结构
        1. 一维：
            1. 基础：数组 array (string), 链表 linked list
            2. 高级：栈 stack, 队列 queue, 双端队列 deque, 集合 set, 映射 map (hash or map), etc
        2. 二维：
            1. 基础：树 tree, 图 graph
            2. 高级：二叉搜索树 binary search tree (red-black tree, AVL), 堆 heap, 并查集 disjoint set, 字典树 Trie, etc
        3. 特殊：
            1. 位运算 Bitwise, 布隆过滤器 BloomFilter
            2. LRU Cache
### 刻意练习deliberate practicing
1. 过遍数，五毒神掌
2. 练弱项
### 反馈feedback
1. 即时反馈
2. 主动型反馈（自己去找）
3. 被动式反馈（高手给你指点）

## 切题四件套
1. clarification审题，和面试官讨论，同事、朋友；
2. possible solutions所有可能解法
    1. compare(time/space)
    2. optimal(加强)
3. coding多写
4. test cases

## 五遍刷题法
1. 5分钟
2. 看解法，然后默写
3. 隔天复习
4. 隔周重复练习，记忆曲线
5. 复习

## tips
1. 时间复杂度七个等级；空间复杂度
2. 主定理
3. 数据结构：一维、二维、特殊
4. 算法：branch,iteration,recursion,search(深度、广度),动态规划,二分查找,贪心，数学+几何
## 数组
1. 做数组题目时，常见解题方法有：双指针法（两端夹逼，快慢），递归，迭代，搜索时可以考虑先排序
## 链表
1. 解法相对固定，刻意练习
2. 检测循环链表，快慢指针
## skip list
1. 有序的链表
2. 链表基础上增加了维度，实现二分查找
3. 优点：原理简单、容易实现
4. 缺点：维护成本高，增加删除（logN）会改变维度
## queue、stack
1. queue,dequeue(double end),priority queue
2. stack,minStack,最近相关性（如括号匹配）
3. 固定窗口的题目，用dequeue解决
