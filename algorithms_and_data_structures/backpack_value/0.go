package main

import (
	"fmt"
	"math"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// f(0, 0, 0)
	// spew.Dump(max)

	f2()
	spew.Dump(max2)

	// f3()
	// spew.Dump(max3)
}

var max int = math.MinInt32
var max2 int = math.MinInt32
var max3 int = math.MinInt32
var items []int = []int{2, 2, 4, 6, 3}
var values []int = []int{3, 4, 8, 9, 6}
var n int = 5
var w int = 9

func f(i int, cw int, cv int) {
	if cw == w || i == n {
		if cv > max {
			max = cv
		}
		return
	}

	f(i+1, cw, cv)        // 不装
	if cw+items[i] <= w { // 装第i个
		f(i+1, cw+items[i], cv+values[i])
	}
}

func f2() {
	states := [5][10]int{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}

	states[0][0] = 0
	if items[0] <= w {
		states[0][items[0]] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-items[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + values[i]
				if v > states[i][j+items[i]] {
					states[i][j+items[i]] = v
				}
			}
		}
	}

	fmt.Println(" w:0  1  2  3  4  5  6  7  8  9")
	fmt.Println("---------------------")
	for k, v := range states {
		fmt.Printf("%v| ", k)
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	for j := w; j >= 0; j-- {
		if states[n-1][j] != -1 {
			max2 = states[n-1][j]
			return
		}
	}
}

func f3() {
	states := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	states[0] = 0
	if items[0] <= w {
		states[items[0]] = values[0]
	}

	for _, v := range states {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for i := 1; i < n; i++ {
		for j := w - items[i]; j >= 0; j-- {
			if states[j] >= 0 {
				v := states[j] + values[i]
				if v > states[j+items[i]] {
					states[j+items[i]] = v
				}
			}
		}
		for _, v := range states {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}

	for _, v := range states {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	for j := w; j >= 0; j-- {
		if states[j] != -1 {
			max3 = states[j]
			return
		}
	}
}
