package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	f(arr, 3, 3)
}

func f(arr []int, n int, k int) {
	// fmt.Printf("%v %v %v\n", arr, n, k)
	if k == 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("%v ", arr[i])
		}
		fmt.Println()
	}

	for i := 0; i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]
		f(arr, n, k-1)
		arr[i], arr[k-1] = arr[k-1], arr[i]
	}
}
