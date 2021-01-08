package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
	. "github.com/9d77v/leetcode/pkg/algorithm/stack"
)

/*
题目：
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
*/

/*
方法一： 栈
时间复杂度：О(n+k)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.4 MB
*/
func getKthFromEndFunc1(head *ListNode, k int) *ListNode {
	var stack Stack = NewSliceStack(0)
	for head != nil {
		stack.Push(head)
		head = head.Next
	}
	for i := 0; i < k; i++ {
		value := stack.Pop()
		if i == k-1 {
			return value.(*ListNode)
		}
	}
	return nil
}

/*
方法二： 快慢双指针
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.2 MB
*/
func getKthFromEndFunc2(head *ListNode, k int) *ListNode {
	fast, slow, i := head, head, 0
	for fast != nil {
		if i < k-1 {
			fast = fast.Next
			i++
		} else {
			if fast.Next != nil {
				slow = slow.Next
			}
			fast = fast.Next
		}
	}
	if i >= k-1 {
		return slow
	}
	return nil
}
