package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/complexlist"
)

/*
题目：复杂链表的复制

请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
*/
/*
运行时间：0 ms	内存消耗：3.4 MB
*/
func copyRandomList(head *Node) *Node {
	cloneNode(head)
	cloneRandom(head)
	return reconnect(head)
}

func cloneNode(head *Node) {
	cur := head
	for cur != nil {
		clonedNode := new(Node)
		clonedNode.Val = cur.Val
		clonedNode.Next = cur.Next
		cur.Next = clonedNode
		cur = clonedNode.Next
	}
}

func cloneRandom(head *Node) {
	cur := head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
}

func reconnect(head *Node) *Node {
	cur := head
	newHead := new(Node)
	newCur := newHead
	for cur != nil {
		newCur.Next = cur.Next
		newCur = newCur.Next
		cur.Next = cur.Next.Next
		cur = cur.Next
	}
	return newHead.Next
}
