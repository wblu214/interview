package main

import "fmt"

func findPublicListNode(l1, l2 *ListNode) *ListNode {
	pre1, pre2 := l1, l2
	var fast, slow *ListNode
	Len1, Len2 := 0, 0
	step := 0

	for pre1 != nil {
		pre1 = pre1.Next
		Len1++
	}
	for pre2 != nil {
		pre2 = pre2.Next
		Len2++
	}
	if Len1 > Len2 {
		step = Len1 - Len2
		fast, slow = l1, l2
	} else {
		step = Len2 - Len1
		fast, slow = l2, l1
	}

	fmt.Println("链表长度差为:", step)
	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != nil && slow != nil {
		if fast.Val != slow.Val {
			fast = fast.Next
			slow = slow.Next
		} else {
			return fast
		}
	}
	return nil

}
func main() {
	l1 := CreateLinkedNode([]int{2, 3, 4})
	l2 := CreateLinkedNode([]int{9, 9, 2, 5, 2, 2, 3, 4})

	PrintLinkedNode(l1)
	PrintLinkedNode(l2)

	resNode := findPublicListNode(l1, l2)
	if resNode != nil {
		fmt.Println("交叉节点为：", resNode.Val)
	} else {
		fmt.Println("没有节点交叉！！！！")

	}
}
