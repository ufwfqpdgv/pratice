package main

import "github.com/davecgh/go-spew/spew"

func main() {
	items := []int{7, 2, 2, 2, 2}
	f(0, 0, items, len(items), 10)
	spew.Dump(max)
}

var max int

func f(i int, cw int, items []int, n int, w int) {
	if cw == w || i == n {
		if cw > max {
			max = cw
		}
		return
	}

	f(i+1, cw, items, n, w)
	if cw+items[i] <= w {
		f(i+1, cw+items[i], items, n, w)
	}
}
