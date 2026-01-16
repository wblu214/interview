package main

import "fmt"

/*
题目描述：
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点

输入：head = [4,5,1,9], val = 5
输出：[4,1,9]
*/

func deleteLinkedNode(head *ListNode, target int) *ListNode {
	fmt.Println("要删除的节点为:", target)
	if head.Val == target {
		return head.Next
	}
	pre := head
	for pre.Next != nil && pre.Next.Val != target {
		pre = pre.Next
	}
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return head
}
func deleteLinkedNode2(head *ListNode, target int) *ListNode {
	fmt.Println("要删除的节点为:", target)
	if head.Val == target {
		return head.Next
	}
	pre, cur := head, head
	for cur != nil && cur.Val != target {
		pre = cur
		cur = cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}
func main() {
	source := CreateLinkedNode([]int{1, 3, 5, 7, 8, 9})
	PrintLinkedNode(source)
	//delLinkedList := deleteLinkedNode(source, 0)
	delLinkedList2 := deleteLinkedNode2(source, 3)

	//printLinkedNode(delLinkedList)
	PrintLinkedNode(delLinkedList2)

}
