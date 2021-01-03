package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：
给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。
0
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
*/

/*
方法一: 模拟
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.4 MB
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
