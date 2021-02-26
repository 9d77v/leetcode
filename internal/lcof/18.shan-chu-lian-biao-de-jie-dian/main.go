package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：删除链表的节点
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

说明：

题目保证链表中节点的值互不相同
若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
*/

/*
方法一：遍历
运行时间：0 ms	内存消耗：2.8 MB
*/
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	prev := head
	if prev.Val == val {
		return head.Next
	}
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			return head
		}
		prev, cur = cur, cur.Next
	}
	return head
}
