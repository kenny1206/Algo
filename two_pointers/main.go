package two_pointers

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	var ans = make([]int, n)
	idx := n-1
	if a == 0 {
		for i := n-1; i >= 0; i-- {
			ans[idx] = fx(nums[i], a, b, c)
			idx--
		}
		if b < 0 {
			reverse(ans)
		}
		return ans
	}
	mid := -1*float32(b)/2/float32(a)
	left, right := 0, len(nums) - 1
	for left <= right {
		if mid - float32(nums[left]) >= float32(nums[right]) - mid {
			ans[idx] = fx(nums[left], a, b, c)
			left++
		} else {
			ans[idx] = fx(nums[right], a, b, c)
			right--
		}
		idx--
	}
	if a < 0 {
		reverse(ans)
	}
	return ans
}

func fx(x int, a int, b int, c int) int {
	return a * x * x + b * x + c
}

func reverse(nums []int) {
	n := len(nums)
	for i:=0; i < n / 2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}