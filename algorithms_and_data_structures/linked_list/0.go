package main

import (
	"fmt"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

// 单链表反转
// 链表中环的检测--同haed起点，一个一次前进一步，另一前进二步，如结束前二者相等，那就有，如无则无
// 两个有序的链表合并
// 删除链表倒数第 n 个结点
// 求链表的中间结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	testReverSeList()
	testMergeList()
	testDelBottomNNode()
	testGetMiddleNode()

	return
}

func reverseList(head *ListNode) (rspHead *ListNode) {
	// 加个哨兵，就不用处理head的情况，方便和容易理解很多
	var a *ListNode
	b := head
	for {
		if b == nil {
			rspHead = a
			return
		}
		bNext := b.Next
		b.Next = a
		a = b
		b = bNext
	}
}

func mergeList(head1 *ListNode, head2 *ListNode) (rspHead *ListNode) {
	// 加个哨兵，就不用处理head的情况，方便和容易理解很多
	tempNode := &ListNode{}
	rspHead = tempNode

	a := head1
	b := head2
	curNode := rspHead
	for {
		if a == nil {
			curNode.Next = b
			break
		} else if b == nil {
			curNode.Next = a
			break
		}

		if a.Val < b.Val {
			curNode.Next = a
			curNode = a
			a = a.Next
		} else {
			curNode.Next = b
			curNode = b
			b = b.Next
		}
	}

	rspHead = rspHead.Next
	return
}

// 递归解法
func mergeList_2(head1 *ListNode, head2 *ListNode) (rspHead *ListNode) {
	if head1 == nil {
		rspHead = head2
		return
	} else if head2 == nil {
		rspHead = head1
		return
	}

	if head1.Val < head2.Val {
		head1.Next = mergeList_2(head1.Next, head2)
		return head1
	} else {
		head2.Next = mergeList_2(head1, head2.Next)
		return head2
	}
}

func delBottomNNode(head *ListNode, n int) (rspHead *ListNode) {
	// 加个哨兵，就不用处理head的情况，方便和容易理解很多
	tempNode := &ListNode{
		Val:  0,
		Next: head,
	}
	rspHead = tempNode

	first := rspHead
	second := rspHead
	for i := 0; i <= n; i++ { // 中间隔n个，即second在n+1的位置
		second = second.Next
	}
	for {
		if second == nil {
			break
		}
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next

	rspHead = rspHead.Next
	return
}

func getMiddleNode(head *ListNode) (rspMiddleNode *ListNode) {
	// 加个哨兵，就不用处理head的情况，方便和容易理解很多
	tempNode := &ListNode{
		Val:  0,
		Next: head,
	}
	first := tempNode
	second := tempNode

	for {
		first = first.Next
		second = second.Next.Next

		if second == nil {
			rspMiddleNode = first
			return
		} else if second.Next == nil {
			rspMiddleNode = first.Next
			return
		}
	}
}

func testReverSeList() {
	var head *ListNode
	var curNode *ListNode

	for i := 1; i <= 5; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == 1 {
			head = tempNode
			curNode = tempNode
		} else {
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	printfList(reverseList(head))
}

func testMergeList() {
	var head1 *ListNode
	var head2 *ListNode
	var curNode *ListNode

	for i := 1; i <= 5; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == 1 {
			head1 = tempNode
			curNode = tempNode
		} else {
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	for i := 3; i <= 6; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == 3 {
			head2 = tempNode
			curNode = tempNode
		} else {
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	// printfList(mergeList(head1, head2))
	printfList(mergeList_2(head1, head2))
}

func testDelBottomNNode() {
	var head *ListNode
	var curNode *ListNode

	for i := 1; i <= 5; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == 1 {
			head = tempNode
			curNode = tempNode
		} else {
			curNode.Next = tempNode
			curNode = tempNode
		}
	}
	printfList(delBottomNNode(head, 4))
}

func printfList(rqHead *ListNode) {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Println("NowFunc:" + runtime.FuncForPC(pc).Name() + " ")

	tempNode := rqHead
	fmt.Println("---------------")
	for {
		if tempNode == nil {
			break
		}
		fmt.Println(tempNode.Val)
		tempNode = tempNode.Next
	}
	fmt.Println("---------------")
}

func testGetMiddleNode() {
	var head *ListNode
	var curNode *ListNode

	for i := 1; i <= 5; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		if i == 1 {
			head = tempNode
			curNode = tempNode
		} else {
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("NowFunc:" + runtime.FuncForPC(pc).Name() + " ")
	spew.Dump(getMiddleNode(head).Val)
}
