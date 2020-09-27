package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{2, &Node{Val: 4}, &Node{Val: 5}}
	root.Right = &Node{3, &Node{Val: 6}, &Node{Val: 7}}
	// middleTraverse(root)
	// breadthTraverse(root)
	depthTraverse(root)
}

func search(root *Node, v int) (rspNode *Node) {
	if root == nil {
		return
	}

	p := root
	for p != nil {
		if v == p.Val {
			rspNode = p
			return
		} else if v < p.Val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	rspNode = nil

	return
}

func add(root *Node, v int) (rspRoot *Node) {
	if root == nil {
		rspRoot = &Node{Val: v}
	}

	p := root
	for p != nil {
		if v < p.Val {
			if p.Left == nil {
				p.Left = &Node{Val: v}
				return
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &Node{Val: v}
				return
			} else {
				p = p.Right
			}
		}
	}

	return
}

func del(root *Node, v int) (rspRoot *Node) {
	if root == nil {
		return
	}

	p := root
	var pFather *Node
	for p != nil && p.Val != v {
		pFather = p
		if v < p.Val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil { // 没找到
		return
	}

	// 待删节点：根节点
	if pFather == nil {
		rspRoot = nil
		return
	}

	// 待删节点：无子节点
	if p.Left == nil && p.Left == nil {
		if pFather.Left == p {
			pFather.Left = nil
		} else {
			pFather.Right = nil
		}
	}

	// 待删节点：仅1个子节点
	if p.Left != nil && p.Right == nil {
		if pFather.Left == p {
			pFather.Left = p.Left
		} else {
			pFather.Right = p.Left
		}
	} else if p.Right != nil && p.Left == nil {
		if pFather.Left == p {
			pFather.Left = p.Right
		} else {
			pFather.Right = p.Right
		}
	}

	// 待删节点：2个子节点，需找到右树最小节点，然后互换值，再删除最小节点
	if p.Left != nil && p.Right != nil {
		q := p.Right
		qFather := p
		for q.Left != nil {
			qFather = q
			q = q.Left
		}

		p.Val = q.Val

		var child *Node
		if q.Right != nil {
			child = q.Right
		}

		if qFather.Left == q {
			qFather.Left = child
		} else {
			qFather.Right = child
		}
	}

	return
}

func del2(root *Node, v int) (rspRoot *Node) {
	if root == nil {
		return
	}

	p := root
	var pFather *Node
	for p != nil && p.Val != v {
		pFather = p
		if v < p.Val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil { // 没找到
		return
	}

	// 待删节点：2个子节点，需找到右树最小节点，然后互换值，再删除最小节点
	if p.Left != nil && p.Right != nil {
		q := p.Right
		qFather := p
		for q.Left != nil {
			qFather = q
			q = q.Left
		}

		p.Val = q.Val
		p = q
		pFather = qFather
	}

	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}

	if pFather == nil {
		rspRoot = nil
		return
	} else if pFather.Left == p {
		pFather.Left = child
	} else {
		pFather.Right = child
	}

	return
}

func middleTraverse(node *Node) {
	if node == nil {
		return
	}

	middleTraverse(node.Left)
	fmt.Printf("%v ", node.Val)
	middleTraverse(node.Right)
}

// 队列（先进先出）
func breadthTraverse(root *Node) {
	l := list.New()
	l.PushBack(root)
	var curNode *Node
	for l.Len() > 0 {
		front := l.Front()
		fmt.Printf("%v ", front.Value.(*Node).Val)
		curNode = front.Value.(*Node)
		l.Remove(front)

		if curNode.Left != nil { //先遍历左树，so先将左树入队列
			l.PushBack(curNode.Left)
		}
		if curNode.Right != nil {
			l.PushBack(curNode.Right)
		}
	}

	return
}

// 栈（先进后出）
func depthTraverse(root *Node) {
	s := &Stack{
		list: &list.List{},
		lock: &sync.RWMutex{},
	}
	s.Push(root)
	var curNode *Node
	for s.Len() > 0 {
		n := s.Peak()
		fmt.Printf("%v ", n.(*Node).Val)
		curNode = n.(*Node)
		s.Pop()

		if curNode.Right != nil { //先遍历左树，so先将右树入栈
			s.Push(curNode.Right)
		}
		if curNode.Left != nil {
			s.Push(curNode.Left)
		}
	}

	return
}

type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	l := &sync.RWMutex{}
	return &Stack{list, l}
}

func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
