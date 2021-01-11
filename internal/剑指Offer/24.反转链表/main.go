package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
	. "github.com/9d77v/leetcode/pkg/algorithm/stack"
)

/*
题目：
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
*/

/*
方法一： 栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3 MB
*/
func reverseListFunc1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var stack Stack = NewSliceStack(0)
	for head != nil {
		stack.Push(head)
		head = head.Next
	}
	reverseHead := stack.Pop().(*ListNode)
	p := reverseHead
	for !stack.Empty() {
		p.Next = stack.Pop().(*ListNode)
		p = p.Next
	}
	p.Next = nil
	return reverseHead
}

/*
方法二： 双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.5 MB
*/
func reverseListFunc2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
