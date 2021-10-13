package merge_intervals

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int
	for len(firstList) > 0 && len(secondList) > 0 {
		f := firstList[0]
		s := secondList[0]
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
	return ans
}

func merge(r1, r2 []int) []int {
	s := max(r1[0], r2[0])
	e := min(r1[1], r2[1])
	if s > e {
		return nil
	}
	return []int{s, e}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}