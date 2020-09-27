package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	arr := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}

	// bubbleSort(arr)
	// insertSort(arr)
	// selectSort(arr)
	// hillSort(arr)
	// arr = MergeSort(arr)
	// quickSortMain(arr)
	// kBig := bigNValue(arr, 6)
	// spew.Dump(kBig, arr[kBig])

	spew.Dump(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for a := 0; a < n-1; a++ {
		exchange := false
		for b := 0; b < n-a-1; b++ {
			if arr[b] > arr[b+1] {
				arr[b], arr[b+1] = arr[b+1], arr[b]
				exchange = true
			}
		}
		if !exchange {
			break
		}
	}
}

func insertSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for a := 1; a < n; a++ {
		v := arr[a]
		b := a - 1
		for ; b >= 0; b-- {
			if arr[b] > v {
				arr[b+1] = arr[b]
			} else {
				break
			}
		}
		arr[b+1] = v
	}
}

func selectSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for a := 0; a < n-1; a++ {
		minIndex := a
		b := a
		for ; b < n; b++ {
			if arr[b] < arr[minIndex] {
				minIndex = b
			}
		}
		if minIndex != a {
			arr[minIndex], arr[a] = arr[a], arr[minIndex]
		}
	}
}

func hillSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for gap := n / 2; gap > 0; gap = gap / 2 {
		for a := gap; a < n; a++ {
			for b := a; b >= gap && arr[b-gap] > arr[b]; b = b - gap { // 因前面已是有序了，所以arr[b-gap] > arr[b]也得满足才再次交换
				arr[b], arr[b-gap] = arr[b-gap], arr[b]
			}
		}
	}
}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	mid := n / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	newArr := []int{}
	l, r := 0, 0
	leftLen := len(left)
	rightLen := len(right)
	for {
		if left[l] < right[r] {
			newArr = append(newArr, left[l])
			l++
			if l == leftLen {
				newArr = append(newArr, right[r:]...)
				break
			}
		} else {
			newArr = append(newArr, right[r])
			r++
			if r == rightLen {
				newArr = append(newArr, left[l:]...)
				break
			}
		}
	}
	fmt.Println(left, right, newArr)
	return newArr
}

func quickSortMain(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	i := l
	j := r
	v := arr[i]
	for i < j {
		for i < j && arr[j] >= v {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] <= v {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = v
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
}

func bigNValue(arr []int, k int) (pos int) {
	pos = quickSort2(arr, 0, len(arr)-1, k-1)

	return
}

func quickSort2(arr []int, l int, r int, k int) (pos int) {
	if l >= r {
		pos = -1
		return
	}

	i := l
	j := r
	v := arr[i]
	for i < j {
		for i < j && arr[j] >= v {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] <= v {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	if i == k {
		pos = i
		return
	} else if i > k {
		pos = quickSort2(arr, l, i-1, k)
	} else if i < k {
		pos = quickSort2(arr, i+1, r, k)
	}

	return
}
