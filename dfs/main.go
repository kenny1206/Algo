package dfs

import "sort"

func minimumTimeRequired(jobs []int, k int) int {
	s := sum(jobs)
	m := max(jobs...)
	left := max(s / k, m)
	right := s
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] > jobs[j]
	})
	for left <= right {
		mid := left + (right - left) / 2
		if canScheduleUnderTime(jobs, mid, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func sum(jobs []int) int {
	var s int
	for _, job := range jobs {
		s += job
	}
	return s
}

func max(jobs ...int) int {
	var s int
	for _, job := range jobs {
		if job > s {
			s = job
		}
	}
	return s
}

func canScheduleUnderTime(jobs []int, time int, k int) bool {
	n := len(jobs)
	if n == 0 {
		return true
	}
	var workers = make([]int, k)
	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if cur == n {
			return true
		}
		curJob := jobs[cur]
		for i := 0; i < k; i++ {
			if workers[i] + curJob > time {
				continue
			}
			workers[i] += curJob
			if dfs(cur+1) {
				return true
			}
			workers[i] -= curJob
			if workers[i] == 0 {
				return false
			}
		}
		return false
	}
	return dfs(0)
}