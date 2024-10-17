### 面试题总结

### 1. DFS与BFS

**1.1 单词搜索**
**难点**
> 基于上下左右计算下一个匹配字符为回溯过程查找有效路径, 单词首字符作为匹配起点
条件
a. 从当前已匹配字符可往上下左右移动
b. 同一个字符不能访问两次

**1.2 有向无环图路径搜索**
**关键点**
> 图的表示方法: 邻接表/邻接矩阵

**1.3 二叉树的层序遍历**
**关键点**
> 基于队列进行层序遍历


**1.4 岛屿问题**
**关键点**
> 理解题意后分而治之, 封装DFS/BFS方法, 控制边界值判断

- [ ] 搜索单词在面板字符序列中是否存在有效路径
- [ ] 图的有效路径
- [ ] 二叉树层序遍历(最小深度和最大深度)

### 2. 二分搜索 

**2.1 寻找两个有序数组的中位数**
难点: 寻找两个有序数组的切割点及边界值判断
算法时间复杂度: O(m+n)

- [ ] 寻找两个有序数组的中位数

### 3. 表达式求值

难点:
1. 一层一层去除括号
2. 判断优先级迭代计算

- [ ]无括号表达式
- [ ]有括号表达式
- [ ]后缀表达式
- [ ]前缀表达式转后缀表达式


### 4. 判断字符串是否为科学计算表示

难点
1. 状态机抽象与转换
2. 分割而治理(https://leetcode.cn/problems/valid-number/solutions/831848/gong-shui-san-xie-zi-fu-chuan-mo-ni-by-a-7cgc/)

### 5. 回文问题

- [回文数组/字符串]
- [回文链表]
- [分割回文字符串](https://www.programmercarl.com/0132.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2II.html#%E6%80%9D%E8%B7%AF)
- [最长回文子串](从下到上,从左到右)

### 6. 最长递增子序列

- [寻找一个最长递增子序列]
- [最长递增子序列的个数]

### 7. 滑动窗口问题

- [] 乘积小于k的子数组个数(https://leetcode.cn/problems/subarray-product-less-than-k/ )
- [] 移动窗口的最大值/最小值

### 8. 哈希表 + 前缀和

- []  0和1数目相等的最长子数组(https://leetcode.cn/problems/contiguous-array/)

### 9. 目标和为0的数据元素

- [x] 目标和为0的3个数(https://leetcode.cn/problems/contiguous-array/)
- [x] 最小大于目标值的子数组(https://leetcode.cn/problems/minimum-size-subarray-sum/description/)
- [x] 有序数组的两个数字和为0(https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/) 