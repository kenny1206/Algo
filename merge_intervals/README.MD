## 题目描述
给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。

返回这 两个区间列表的交集 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/interval-list-intersections

著作权归领扣网络所有。

## 思路
1. 令两个数组的下标的起始值为 i1, i2
2. 计算 firstList[i1] 和 secondList[i2] 的交集
   1. 交集起始值为两个区间的较大起始值，终止值为两个区间的较小终止值
3. 将终止值较小的区间的下标加 1，继续步骤 2
4. 如果任意数组剩余空间为空，则终止

## Template

````
var ans [][]int
for len(firstList) > 0 && len(secondList) > 0 {
   f := firstList[0]
   s := secondList[0]
   // logic for ranges
   cur := merge(f, s)
   if cur != nil {
      ans = append(ans, cur)
   }
   if f[1] < s[1] {
      firstList = firstList[1:]
   } else {
      secondList = secondList[1:]
   }
}
````