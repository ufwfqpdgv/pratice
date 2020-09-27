package main

import "github.com/davecgh/go-spew/spew"

func main() {
	spew.Dump(1)
	spew.Dump(f(4))
}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}
