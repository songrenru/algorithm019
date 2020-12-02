# 学习笔记
## dfs、bfs
1. dfs:
    1. 模板：
    ```python
        def dfs(node):
            if node in visited:
                # already visited
                return
            visited.add(add)
            # process current node
            # ... # logic here
            dfs(node.left)
            dfs(node.right)
    ```
    ```python
        def DFS(self, tree): 

            if tree.root is None: 
                return [] 

            visited, stack = [], [tree.root]

            while stack: 
                node = stack.pop() 
                visited.add(node)

                process (node) 
                nodes = generate_related_nodes(node) 
                stack.push(nodes) 

            # other processing work 
    ```
2. bfs:
    1. 模板：
    ```python
        def bfs(graph, start, end):
            queue = []
            queue.append(start)
            visited.add(start)

            while queue:
                node = queue.popleft() # nodes?
                visited.add(node)

                process(node)
                nodes = generate_related_nodes(node) # newNodes?
                queue.push(nodes) # nodes = newNodes
            # other processing work 
    ```
3. 优先级优先搜索（启发式搜索）
## 贪心greedy算法
1. 贪心：局部最优，以实现全局最优；要点：证明贪心能实现全局最优；
2. 动态规划：保存了过程，有回溯步骤，全局最优；
## 二分查找
1. code template:
    ```python
        left, right = 0, len(arr) - 1
        while left <= right # <=
            mid = left + (right - left) / 2
            if arr[mid] == target:
                break or return
            elif arr[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
    ```