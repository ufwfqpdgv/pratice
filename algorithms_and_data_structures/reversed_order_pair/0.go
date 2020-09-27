package main

import (
	"github.com/davecgh/go-spew/spew"
)

var num int

func main() {
	arr := []int{2, 4, 3, 1, 5, 6}
	spew.Dump(MergeSort(arr))
	spew.Dump(num)
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
		if left[l] <= right[r] {
			newArr = append(newArr, left[l])
			l++
			if l == leftLen {
				newArr = append(newArr, right[r:]...)
				break
			}
		} else {
			num += len(left) - l
			newArr = append(newArr, right[r])
			r++
			if r == rightLen {
				newArr = append(newArr, left[l:]...)
				break
			}
		}
	}
	return newArr
}
