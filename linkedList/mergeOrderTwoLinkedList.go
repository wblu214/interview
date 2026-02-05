package main

/*
*
给定两个以 有序链表 形式记录的训练计划 l1、l2，分别记录了两套核心肌群训练项目编号，请合并这两个训练计划，按训练项目编号 升序 记录于链表并返回。

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func mergeTwoLinkedList(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2

	head := new(ListNode)
	target := head

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			target.Next = cur1
			cur1 = cur1.Next
		} else {
			target.Next = cur2
			cur2 = cur2.Next
		}
		target = target.Next
	}
	//
	if cur1 == nil {
		target.Next = cur2
	}

	if cur2 == nil {
		target.Next = cur1
	}

	return head.Next
}

func main() {
	l1 := CreateLinkedNode([]int{2, 3, 4})
	l2 := CreateLinkedNode([]int{1, 3, 4})

	PrintLinkedNode(l1)
	PrintLinkedNode(l2)

	targetLinkedNode := mergeTwoLinkedList(l1, l2)
	PrintLinkedNode(targetLinkedNode)
}
