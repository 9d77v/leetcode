package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：反转链表
反转一个单链表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list/
*/

/*
方法一：迭代
时间复杂度：О(n)
空间复杂度：O(n)
运行时间：0 ms	内存消耗：2.6 MB
*/
func reverseListFunc1(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		prev = &ListNode{
			Val:  cur.Val,
			Next: prev,
		}
		cur = cur.Next
	}
	return prev
}

/*
方法二：递归
时间复杂度：О(n)
空间复杂度：O(n)
运行时间：0 ms	内存消耗：3.1 MB
*/
func reverseListFunc2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListFunc2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/*
方法三：迭代
时间复杂度：О(n)
空间复杂度：O(1)
运行时间：0 ms	内存消耗：2.5 MB
*/
func reverseListFunc3(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
