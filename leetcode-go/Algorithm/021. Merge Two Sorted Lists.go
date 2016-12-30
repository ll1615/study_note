package main

import (
	"fmt"
	"strconv"
)

func main() {
	l1 := GenerateList([]int{1, 3, 5})
	l1.Print()
	l2 := GenerateList([]int{3, 5, 8, 10})
	l2.Print()
	mergeTwoLists(l1, l2).Print()
}

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//GenerateList aa
func GenerateList(arr []int) *ListNode {
	ln := len(arr)
	if ln == 0 {
		return nil
	}

	var header = &ListNode{Val: arr[0]}
	node := header

	for i := 1; i < ln; i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	return header
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var header, node *ListNode
	node = new(ListNode)
	header = node
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else if l2 != nil {
		node.Next = l2
	}
	return header.Next
}

//Print wt
func (l *ListNode) Print() {
	if l == nil {
		return
	}
	var s string
	for l != nil {
		s += strconv.Itoa(l.Val) + "-->"
		l = l.Next
	}
	fmt.Println(s[:len(s)-3])
}
