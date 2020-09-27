package main

func main() {
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 7, 7, 8}
	// spew.Dump(binarySearch(arr, 7))
	// spew.Dump(binarySearchRecursive(arr, 0, len(arr)-1, 7))
	// spew.Dump(binarySearchFirstEqualValue(arr, 7))
	// spew.Dump(binarySearchLastestEqualValue(arr, 7))

	// arr := []int{1, 3, 5, 7, 9}
	// spew.Dump(arr[binarySearchFirstGreaterEqualValue(arr, 4)])
	// spew.Dump(arr[binarySearchLastestLessEqualValue(arr, 2)])

	// arr := []int{7, 9, 0, 1, 3}
	// spew.Dump(arr[binaryFindMinimumValue(arr)])
}

func binarySearch(arr []int, v int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] == v {
			pos = mid
			return
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	pos = -1
	return
}

func binarySearchRecursive(arr []int, l int, r int, v int) (pos int) {
	if l > r {
		pos = -1
		return
	}

	mid := l + ((r - l) >> 1)
	if arr[mid] == v {
		pos = mid
		return
	} else if arr[mid] > v {
		pos = binarySearchRecursive(arr, l, mid-1, v)
	} else {
		pos = binarySearchRecursive(arr, mid+1, r, v)
	}

	return
}

func binarySearchFirstEqualValue(arr []int, v int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] == v {
			if mid == 0 || arr[mid-1] < 0 {
				pos = mid
				return
			} else {
				r = mid - 1
			}
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	pos = -1
	return
}

func binarySearchLastestEqualValue(arr []int, v int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] == v {
			if mid == len(arr)-1 || arr[mid+1] > v {
				pos = mid
				return
			} else {
				l = mid + 1
			}
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	pos = -1
	return
}

func binarySearchFirstGreaterEqualValue(arr []int, v int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] >= v {
			if mid == 0 || arr[mid-1] < v {
				pos = mid
				return
			} else {
				r = mid - 1
			}
		} else {
			l = mid + 1
		}
	}
	return
}

func binarySearchLastestLessEqualValue(arr []int, v int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if arr[mid] <= v {
			if mid == len(arr)-1 || arr[mid+1] > v {
				pos = mid
				return
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}
	return
}

func binaryFindMinimumValue(arr []int) (pos int) {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := l + ((r - l) >> 1)
		if arr[mid] > arr[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	pos = l
	return
}
