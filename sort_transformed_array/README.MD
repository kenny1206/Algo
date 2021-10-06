## 题目描述
Leetcode 360. 有序转化数组

给你一个已经 排好序 的整数数组 nums 和整数 a、b、c。对于数组中的每一个数 x，计算函数值 f(x) = ax2 + bx + c，请将函数值产生的数组返回。

要注意，返回的这个数组必须按照 升序排列，并且我们所期望的解法时间复杂度为 O(n)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-transformed-array

著作权归领扣网络所有。

## 思路
1. 二元一次方程的值，是一个抛物线，但需要考虑 a 和 b 的取值
   1. 若 a > 0, 最小点在 mid=-1*b/2/a, x 离 mid 越远，函数值越大
   2. 如 a < 0, 最大点在 mid=-1*b/2/a, x 离 mid 越远，函数值越小
   3. 若 a == 0 && b >= 0, 函数值单调递增
   4. 若 a == 0 && b < 0, 函数值单调递减
2. 我们可以先按 1.1 或 1.3 计算出 ans 数组，如果是 1.2 或 1.4 的情况，就将 ans 反转

## Template

````
var ans []int
var idx int
for left <= right {
    if condition(left) {
        ans[idx] = fn(left)
        left++
    } else {
        ans[idx] = fn(right)
        right--
    }
    idx++
}
````