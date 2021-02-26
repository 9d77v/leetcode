package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：合并2个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
*/

/*
方法一： 遍历
运行时间：4 ms	内存消耗：4.1 MB
*/
func mergeTwoListsFunc1(l1 *ListNode, l2 *ListNode) *ListNode {
	mergedHead := new(ListNode)
	cur := mergedHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil && l2 != nil {
		cur.Next = l2
	}
	if l2 == nil && l1 != nil {
		cur.Next = l1
	}
	return mergedHead.Next
}

/*
方法二： 递归
运行时间：8 ms	内存消耗：4.3 MB
*/
func mergeTwoListsFunc2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoListsFunc2(l1.Next, l2)
	return l1
}
