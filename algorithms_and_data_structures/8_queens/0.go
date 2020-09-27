package main

import (
	"fmt"
)

var arr [8]int

func main() {
	cal8Queens(0)
}

func cal8Queens(row int) {
	if row == 8 {
		p(arr)
		return
	}

	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			arr[row] = column
			cal8Queens(row + 1)
		}
	}
}

func isOk(row int, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row - 1; i >= 0; i-- {
		if arr[i] == column {
			return false
		}
		if leftUp >= 0 && arr[i] == leftUp {
			return false
		}
		if rightUp < 8 && arr[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}

	return true
}

func p(arr [8]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
