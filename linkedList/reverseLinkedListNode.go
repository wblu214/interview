package main

func reverseLinkedListNode(head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	var newHead *ListNode
	for cur != nil {
		node := cur.Next   // 保存当前节点的下一个节点
		cur.Next = newHead //反转节点指向
		newHead = cur
		cur = node
	}
	PrintLinkedNode(newHead)
}
func main() {
	head := CreateLinkedNode([]int{2, 3, 6, 8, 2, 1, 4, 6, 0, 19})
	PrintLinkedNode(head)
	reverseLinkedListNode(head)
}
