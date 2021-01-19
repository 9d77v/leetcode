package main

import (
	"math"
)

/*
题目：
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
*/

/*
运行时间：0 ms	内存消耗：2 MB
*/
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	n := len(postorder)
	root := postorder[n-1]
	i := 0
	for ; i < n-1; i++ {
		if postorder[i] > root {
			break
		}
	}
	j := i
	for ; j < n-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	left := true
	if i > 0 {
		left = verifyPostorder(postorder[:i])
	}
	right := true
	if i < n-1 {
		right = verifyPostorder(postorder[i : n-1])
	}
	return left && right
}

/*
方法二： 单调栈
时间复杂度：О(n+k)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.4 MB
*/
func verifyPostorderFunc2(postorder []int) bool {
	n := len(postorder)
	if n == 0 {
		return true
	}
	root := math.MaxInt64
	stack := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
