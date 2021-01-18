package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
	. "github.com/9d77v/leetcode/pkg/algorithm/queue"
)

/*
题目：
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
*/

/*
方法一： bfs,双端队列（切片）
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：4 MB
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	var deque Deque = NewListDeque(1)
	deque.PushBack(root)
	for !deque.Empty() {
		//打印奇数层
		tmp := make([]int, 0)
		n := deque.Len()
		for i := 0; i < n; i++ {
			node := deque.PopFront().(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}
		}
		result = append(result, tmp)
		if deque.Empty() {
			break
		}
		//打印偶数层
		tmp = make([]int, 0)
		n = deque.Len()
		for i := 0; i < n; i++ {
			node := deque.PopBack().(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Right != nil {
				deque.PushFront(node.Right)
			}
			if node.Left != nil {
				deque.PushFront(node.Left)
			}
		}
		result = append(result, tmp)
	}
	return result
}

/*
方法二： bfs,双端队列（链表）
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.9 MB
*/
func levelOrderFunc2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	var deque Deque = NewListDeque(1)
	deque.PushBack(root)
	for !deque.Empty() {
		//打印奇数层
		tmp := make([]int, 0)
		n := deque.Len()
		for i := 0; i < n; i++ {
			node := deque.PopFront().(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}
		}
		result = append(result, tmp)
		if deque.Empty() {
			break
		}
		//打印偶数层
		tmp = make([]int, 0)
		n = deque.Len()
		for i := 0; i < n; i++ {
			node := deque.PopBack().(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Right != nil {
				deque.PushFront(node.Right)
			}
			if node.Left != nil {
				deque.PushFront(node.Left)
			}
		}
		result = append(result, tmp)
	}
	return result
}

/*
方法三： 获取结果，反转
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：0 ms	内存消耗：2.8 MB
*/
func levelOrderFunc3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	var queue Queue = NewSliceQueue(1, root)
	level := 0
	for !queue.Empty() {
		n := queue.Len()
		result = append(result, []int{})
		for i := 0; i < n; i++ {
			node := queue.Pop().(*TreeNode)
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		level++
	}
	for i, v := range result {
		if i%2 == 1 {
			Reverse(v)
		}
	}
	return result
}
