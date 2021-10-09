## 题目描述
1723. 完成所有工作的最短时间

给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。

请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。工人的 工作时间 是完成分配给他们的所有工作花费时间的总和。请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs

著作权归领扣网络所有。

## 思路
1. 二分查找最大时间，上界为 sum(jobs), 下界为 max(sum(jobs)/k, max(jobs))
2. 针对二分的最大时间 mid, 定义 workers 数组, dfs 给 workers 分配任务，判断 mid 能否满足要求
3. dfs 耗时需要 k^n, 但我们可以进行适度的优化：
   1. 优先分配耗时大的任务
   2. 上一个 worker 未分配任务时，不需要给下一个工人分配

## Template for DFS

````
dfs = func(cur int) bool {
   // 终止条件，这里也可以剪枝
   if end_condition {
      return true
   }
   // 做选择
   forin choices {
      // 剪枝
      if not condition {
          continue
      }
      // 保存现场
      do choice
      if dfs(cur+1) {
          // 找到解，提前终止
          return true
      }
      // 还原现场，做下一个选择
      reset choice
   }
   return false
}
````