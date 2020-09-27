package main

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	shortest_path(0, 0, 0, len(arr)-1)
	// shortest_path_2(len(arr))
	spew.Dump(minDist)
}

var arr [4][4]int = [4][4]int{{0, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
var minDist int = math.MaxInt32

func shortest_path(i, j, dist, n int) {
	if i == n && j == n {
		if dist < minDist {
			minDist = dist
		}
		return
	}
	if i < n {
		shortest_path(i+1, j, dist+arr[i][j], n)
	}
	if j < n {
		shortest_path(i, j+1, dist+arr[i][j], n)
	}
}

func shortest_path_2(n int) {
	states := [4][4]int{}
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i][0]
		states[i][0] = sum
	}
	sum = 0
	for j := 0; j < n; j++ {
		sum += arr[0][j]
		states[0][j] = sum
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if states[i-1][j] < states[i][j-1] {
				states[i][j] = states[i-1][j] + arr[i][j]
			} else {
				states[i][j] = states[i][j-1] + arr[i][j]
			}
		}
	}

	minDist = states[n-1][n-1]
}
