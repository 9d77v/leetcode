package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/

/*
方法一：逐位相加
时间复杂度：О(max(m,n))
空间复杂度：О(max(m,n))
运行时间：12 ms	内存消耗：4.9 MB
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result, carry := &ListNode{}, 0
	cur := result
	for l1 != nil || l2 != nil || carry != 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		cur.Val = sum % 10
	}
	return result.Next
}
