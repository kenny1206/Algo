package binary_search


// MountainArray This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

func NewMountainArray(arr []int) *MountainArray {
	m := &MountainArray{
		arr: arr,
	}
	return m
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := findPeak(mountainArr, 0, mountainArr.length()-1)
	left := binarySearch(target, mountainArr, 0, peak, true)
	if left != -1 {
		return left
	}
	return binarySearch(target, mountainArr, peak+1, mountainArr.length()-1, false)
}

func findPeak(mountainArr *MountainArray, left, right int) int {
	if left == right {
		return left
	}
	mid := left + (right-left) / 2
	midVal := mountainArr.get(mid)
	nextVal := mountainArr.get(mid+1)
	if nextVal > midVal {
		return findPeak(mountainArr, mid+1, right)
	} else {
		return findPeak(mountainArr, left, mid)
	}
}

func binarySearch(target int, mountainArr *MountainArray, left, right int, incr bool) int {
	if left > right {
		return -1
	}
	mid := left + (right-left) / 2
	midVal := mountainArr.get(mid)
	if midVal == target {
		return mid
	} else if midVal > target {
		if incr {
			return binarySearch(target, mountainArr, left, mid-1, incr)
		} else {
			return binarySearch(target, mountainArr, mid+1, right, incr)
		}
	} else {
		if incr {
			return binarySearch(target, mountainArr, mid+1, right, incr)
		} else {
			return binarySearch(target, mountainArr, left, mid-1, incr)
		}
	}
}
