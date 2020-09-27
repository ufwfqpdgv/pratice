package main

import (
	"fmt"
)

func main() {
	arr := []byte{'D', 'a', 'F', 'B', 'c', 'A', 'z'}
	lowercase_left_uppercase_right(arr)

	for _, v := range arr {
		fmt.Printf("%v ", string(v))
	}

	// 如加入数字则用桶更好；如不用额外空间则先分2大类，再把含2类的大类再内部快排一下
}

func lowercase_left_uppercase_right(arr []byte) {
	if len(arr) <= 1 {
		return
	}

	l := 0
	r := len(arr) - 1
	for l < r {
		for l < r && !('A' <= arr[l] && arr[l] <= 'Z') {
			l++
		}
		for l < r && !('a' <= arr[r] && arr[r] <= 'z') {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	return
}
