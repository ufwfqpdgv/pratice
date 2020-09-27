package main

import (
	"fmt"
)

// 输入（先序）
// abc##de#g##f###
// 输出(中序
// c b e g d f a

type Node struct {
	V byte
	L *Node
	R *Node
}

var input = "abc##de#g##f###"
var root *Node
var index = 0

func main() {
	root = addNode(input)

	preOutPut(root)
	fmt.Println()

	midOutPut(root)
	fmt.Println()

	laterOutPut(root)
	fmt.Println()
}

func addNode(s string) (rspRoot *Node) {
	c := s[index]
	index += 1
	if c == '#' {
		return
	}
	newNode := &Node{V: c}
	rspRoot = newNode
	newNode.L = addNode(s)
	newNode.R = addNode(s)

	return
}

func preOutPut(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%c ", root.V)
	preOutPut(root.L)
	preOutPut(root.R)
}

func midOutPut(root *Node) {
	if root == nil {
		return
	}
	midOutPut(root.L)
	fmt.Printf("%c ", root.V)
	midOutPut(root.R)
}

func laterOutPut(root *Node) {
	if root == nil {
		return
	}
	laterOutPut(root.L)
	laterOutPut(root.R)
	fmt.Printf("%c ", root.V)
}
