package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedNode(sourceArray []int) *ListNode {

	if len(sourceArray) == 0 {
		return nil
	}
	head := &ListNode{Val: sourceArray[0]}
	cur := head

	for i := 1; i < len(sourceArray); i++ {
		node := &ListNode{Val: sourceArray[i], Next: nil}
		cur.Next = node
		cur = node
	}
	return head
}

func PrintLinkedNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
	fmt.Println()
}
