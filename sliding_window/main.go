package sliding_window

type deque struct {
	values []int
	compare func(v1, v2 int) bool
}

// RPush Keep a monotonous stack
func (dq *deque) RPush(v int) {
	for {
		n := len(dq.values)
		if n == 0 || dq.compare(dq.values[n-1], v) {
			break
		}
		dq.values = dq.values[:n-1]
	}
	dq.values = append(dq.values, v)
}

func (dq *deque) First() int {
	if len(dq.values) == 0 {
		return -1
	}
	return dq.values[0]
}

func (dq *deque) LPop() {
	if len(dq.values) == 0 {
		return
	}
	dq.values = dq.values[1:]
}

func NewDeque(cmp func(v1, v2 int) bool) *deque {
	return &deque{
		compare: cmp,
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < k {
		return nil
	}
	var ans []int
	dq := NewDeque(func(v1, v2 int) bool {
		return nums[v1] > nums[v2]
	})
	for i := 0; i < k-1; i++ {
		dq.RPush(i)
	}
	for right := k-1; right < n; right++ {
		dq.RPush(right)
		for dq.First() < right - k + 1 {
			dq.LPop()
		}
		f := dq.First()
		ans = append(ans, nums[f])
	}
	return ans
}