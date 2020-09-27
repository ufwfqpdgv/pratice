package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// arr := []int{6, 4, 3, 1}
	// add(arr, len(arr), 9, 7)
	// delTopElement(arr, len(arr)-1, 0)

	// arr2 := []int{3, 4, 6, 1}
	// buildBigHeap(arr2, len(arr2)-1)
	// spew.Dump(arr2)
	// heapSort(arr2, len(arr2)-1)
}

func add(arr []int, count int, max int, data int) {
	if count >= max {
		return
	}

	arr = append(arr, data)

	for i := count; (i-1)/2 >= 0 && arr[i] > arr[(i-1)/2]; i = (i - 1) / 2 {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
	}

	spew.Dump(arr)

	return
}

func delTopElement(arr []int, lastIndex int, delPos int) {
	arr[delPos] = arr[lastIndex]
	arr[lastIndex] = -1
	heapify(arr, lastIndex, 0)

	spew.Dump(arr)

	return
}

func buildBigHeap(arr []int, lastIndex int) {
	for index := (lastIndex - 1) / 2; index >= 0; index-- {
		heapify(arr, lastIndex, index)
	}

	return
}

func heapify(arr []int, lastIndex int, index int) {
	for {
		i := index
		nextPos := index
		if i*2+1 <= lastIndex && arr[i] < arr[i*2+1] {
			arr[i], arr[i*2+1] = arr[i*2+1], arr[i]
			nextPos = nextPos*2 + 1
		}
		if i*2+2 <= lastIndex && arr[i] < arr[i*2+2] {
			arr[i], arr[i*2+2] = arr[i*2+2], arr[i]
			nextPos = nextPos*2 + 2
		}

		if nextPos == i {
			break
		}
		i = nextPos
	}
}

func heapSort(arr []int, lastIndex int) []int {
	buildBigHeap(arr, len(arr)-1)
	for index := lastIndex; index > 0; index-- {
		arr[index], arr[0] = arr[0], arr[index]
		heapify(arr, index-1, 0)
	}
	spew.Dump(arr)

	return arr
}
