package main

import "github.com/davecgh/go-spew/spew"

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{1, 2, 6}
	rspArr := mergeTwoOrderlyArr(arr1, arr2)
	spew.Dump(rspArr)
}

func mergeTwoOrderlyArr(arr1 []int, arr2 []int) (rspArr []int) {
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			rspArr = append(rspArr, arr1[i])
			i++
		} else {
			rspArr = append(rspArr, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		rspArr = append(rspArr, arr1[i:]...)
	} else if j < len(arr2) {
		rspArr = append(rspArr, arr2[j:]...)
	}

	return
}
