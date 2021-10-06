## 题目描述
Leetcode 239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum

著作权归领扣网络所有。
## 思路
1. 收集窗口 [0,k-1) 的信息
2. right 从 k-1 开始，将 right 加入窗口
3. 判断是否要将 left 移出窗口
4. 取当前窗口的值

## Template

````
for right := k-1; right < n; right++ {
    addRight()
    removeLeft()
    calculateWindow()
}
````