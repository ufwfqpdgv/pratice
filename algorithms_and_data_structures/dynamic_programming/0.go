package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	weightArr := []int{2, 2, 4, 4, 3} // n物品求最大可组出来的值
	spew.Dump(f(weightArr, len(weightArr), 9))
	// spew.Dump(f2(weightArr, len(weightArr), 9))
}

func f(weight []int, n int, w int) int {
	states := [5][10]bool{}

	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}

	fmt.Println(" w:0 1 2 3 4 5 6 7 8 9")
	fmt.Println("---------------------")
	for k, v := range states {
		fmt.Printf("%v| ", k)
		for _, v2 := range v {
			if v2 {
				fmt.Printf("%v ", 1)
			} else {
				fmt.Printf("%v ", 0)
			}
		}
		fmt.Println()
	}

	// 输出>=8 && <=9的最小值、及其一条可行路径
	j := 0
	for j = 8; j <= 9; j++ {
		if states[n-1][j] {
			break
		}
	}
	if j > 9 {
		spew.Dump("不存在// 输出>=8 && <=9的最小值、及其可能路径")
	} else {
		// // 存在，找所有可能的路径
		ddd(n-1, j, states, weight, []int{})
		fmt.Println()

		// 存在，找一条可达到8的路径
		i := 0
		for i = n - 1; i >= 1; i-- {
			if j-weight[i] >= 0 && states[i-1][j-weight[i]] {
				fmt.Printf("%v ", weight[i])
				j = j - weight[i]
			}
		}
		if j != 0 { // 第0个也选择了，所以j没归0
			fmt.Printf("%v ", weight[0])
		}
		fmt.Println()
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}

	return -1
}

func ddd(i int, j int, states [5][10]bool, weight []int, arr2 []int) {
	if i == 0 {
		if j != 0 { // 第0个也选择了，所以j没归0
			arr2 = append(arr2, 0)
		}
		for _, v := range arr2 {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
		return
	}

	if i-1 >= 0 && j >= 0 && states[i-1][j] {
		ddd(i-1, j, states, weight, arr2)
	}

	if i-1 >= 0 && j-weight[i] >= 0 && states[i-1][j-weight[i]] {
		arr2 = append(arr2, i)
		ddd(i-1, j-weight[i], states, weight, arr2)
	}
}

func f2(weight []int, n int, w int) int {
	states := [10]bool{}

	states[0] = true
	if weight[0] <= w {
		states[weight[0]] = true
	}

	for _, v := range states {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for i := 1; i < n; i++ {
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] {
				states[j+weight[i]] = true
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

	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}

	return -1
}
