# 学习笔记
## 字典树和并查集
### 字典树
1. 定义：字典树，即 Trie 树，又称单词查找树或键树，是一种树形结构。典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
1. 优点：最大限度地减少无谓的字符串比较，查询效率比哈希表高。
### 并查集Disjoint Set（算法4的联通分量）
1. 适用场景：组团、配对问题
2. 基本操作：
    1. makeSet(s):  建立一个新的并查集，其中包含s个单元素集合；
    2. unionSet(x, y):  把元素x和元素y所在的集合合并，要求x和y所在的集合不想交，如果相交则不合并；
    3. find(x): 找到元素x所在的集合的代表，该操作可用于判断两个元素是否位于同一个集合；
3. 路径压缩
4. [并查集代码模板](https://shimo.im/docs/VtcxL0kyp04OBHak/read)
## 高级搜索
### 剪枝
### 双向bfs
### 启发式搜索Heuristic Search(A*)
1. [A* 代码模板](https://shimo.im/docs/8CzMlrcvbWwFXA8r/read)
#### 估价函数
1. 启发式函数：h(n),它用来评价哪些结点最有希望的 是我们想要找到结点，h(n)会返回一个非负实数，也可以认为是从结点n的目标结点路径的估计成本
2. 启发式函数是一种告知搜索方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标。
## AVL树、红黑树
### AVL树
1. 平衡二叉搜索树
2. Balance Factor（平衡因子）：是它的左子树的高度减去它的右子树的高度（有时相反）。balance factor = {-1, 0, 1}
3. 通过旋转操作来进行平衡（四种）
4. 不足：结点需要存储额外信息、且调整次数频繁
### 红黑树
1. 红黑树是一种**近似平衡**的**二叉搜索树**（Binary Search Tree），它够确保任何一个结点的左右子树的高度差小于**两倍**。
    1. 根结点是黑色
    2. 每个叶结点（NIL结点，空结点）是黑色的
    3. 不能有相邻接的两个红色结点
    4. 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点
### 对比
1. AVL trees provide **faster lookups** than Red Black Trees because they are **more strictly balanced**
2. Red Black Trees provide **faster insertion and removal operations** than AVL trees as fewer rotations are done due to relatively relaxed balancing.
3. AVL trees storebalance **factors or heights** with each node, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of information per node.
4. Red Black Trees are used in most of the **language libraries** like map, multimap, multisetin C++ whereas AVL trees are used in **databases** where faster retrievals are required