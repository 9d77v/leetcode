package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

/*
题目：从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
限制：
0 <= 链表长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
*/

/*
方法一：递归
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：4.7 MB
*/
func reversePrintFunc1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	res := reversePrintFunc1(head.Next)
	return append(res, head.Val)
}

/*
方法二：栈,顺序存数组再反转数组
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.1 MB
*/
func reversePrintFunc2(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}

/*
方法三：栈，获取链表长度，反向赋值数组
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.8 MB
*/
func reversePrintFunc3(head *ListNode) []int {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	res := make([]int, n)
	for ; head != nil; head = head.Next {
		res[n-1] = head.Val
		n--
	}
	return res
}
