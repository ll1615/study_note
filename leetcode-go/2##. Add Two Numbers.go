package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := []int{2, 4, 3}
	n2 := []int{5, 6, 4}

	l1, l2 := initList(n1), initList(n2)

	printList(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  //first way
	tail := new(ListNode)
	list := tail
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmp := val1 + val2 + carry

		if tmp >= 10 {
			tail.Val = tmp % 10
			carry = 1
		} else {
			tail.Val = tmp
			carry = 0
		}

		if !(l1 == nil && l2 == nil && carry == 0){
			tail.Next = new(ListNode)
			tail = tail.Next
		}
	}
	return list
	
	//second way
	//tail := new(ListNode)
	//list := tail
	//carry := 0
	//for l1 != nil || l2 != nil || carry > 0 {
	//	node := new(ListNode)
	//
	//	val1, val2 := 0, 0
	//	if l1 != nil {
	//		val1 = l1.Val
	//		l1 = l1.Next
	//	}
	//	if l2 != nil {
	//		val2 = l2.Val
	//		l2 = l2.Next
	//	}
	//
	//	tmp := val1 + val2 + carry
	//	node.Val = tmp % 10
	//	carry = tmp / 10
	//
	//	tail.Next = node
	//	tail = node
	//}
	//return list.Next
}

func printList(list *ListNode) {
	output := "["
	for list != nil {
		output += fmt.Sprint(list.Val) + ","
		list = list.Next
	}
	output = output[:len(output)-1] + "]"
	fmt.Println(output)
}

//generate *ListNode for an int array
func initList(nums []int) *ListNode {
	length := len(nums)
	listTail:= new(ListNode)
	list := listTail

	for i := 0; i < length; i++ {
		node := new(ListNode)
		node.Val = nums[i]
		listTail.Next = node
		listTail = node
	}

	return list.Next
}
