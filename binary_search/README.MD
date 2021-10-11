## 题目描述
给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。

如果不存在这样的下标 index，就请返回 -1。

何为山脉数组？如果数组 A 是一个山脉数组的话，那它满足如下条件：
````
首先，A.length >= 3

其次，在 0 < i < A.length - 1 条件下，存在 i 使得：

A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]
````
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-in-mountain-array

著作权归领扣网络所有。

## 思路
1. 找出[left, right] 区间的山峰 peak，
   1. 如果 left == right, 则 peak = left。否则计算 mid 和 mid+1 的值，计为 value 和 value1，
   2. 如果 value1 > value, 则 peak 在 [mid+1, right]
   3. 如果 value1 < valu, 则 peak 在 [left, mid-1]
2. 两个区间分别二分查找
3. 这里用递归简化代码，二分通常会使用迭代，可以节省栈空间

## Template for BinarySearch

````
func binarySearch(target int, arr []int) int {
   left, right := 0, len(arr) - 1
   for left <= right {
      mid := left + (right - left) / 2
      if arr[mid] == target {
         return mid
      } else if arr[mid] > target {
         right = mid - 1
      } else if arr[mid] < target {
         left = mid + 1
      }
   }
   return -1
}
````