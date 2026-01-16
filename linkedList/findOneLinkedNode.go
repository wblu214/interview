package main

import "fmt"

/*
题目详情：
给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。

输入：head = [2,4,7,8], cnt = 1
输出：8
*/

func findKNodeReverseOrder(head *ListNode, count int) *ListNode {
	quick, slow := head, head

	for quick.Next != nil && count > 0 {
		quick = quick.Next
		count--
	}
	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}
func main() {

	head := CreateLinkedNode([]int{2, 3, 6, 8, 2, 1, 4, 6, 0, 19})
	PrintLinkedNode(head)
	targetLinkedNode := findKNodeReverseOrder(head, 0)
	fmt.Println(targetLinkedNode.Val)
}
