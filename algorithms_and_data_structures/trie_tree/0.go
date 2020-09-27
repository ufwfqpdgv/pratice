package main

import (
	"github.com/davecgh/go-spew/spew"
)

var root *Node

func main() {
	root = &Node{Data: '/'}
	insert("abc")
	insert("abd")
	spew.Dump(root)
	spew.Dump(find("ab"))
	spew.Dump(find("abc"))
	spew.Dump(find("abs"))
}

type Node struct {
	Data      byte
	ChildNode [26]*Node
	IsEnd     bool
}

func insert(rqStr string) {
	p := root
	for _, v := range []byte(rqStr) {
		pos := v - 'a'
		if p.ChildNode[pos] == nil {
			p.ChildNode[pos] = &Node{Data: v}
		}
		p = p.ChildNode[pos]
	}
	p.IsEnd = true
}

func find(rqStr string) bool {
	p := root

	for _, v := range []byte(rqStr) {
		pos := v - 'a'
		if p.ChildNode[pos] == nil {
			return false
		}
		p = p.ChildNode[pos]
	}

	if p.IsEnd {
		return true
	} else {
		return false
	}
}
