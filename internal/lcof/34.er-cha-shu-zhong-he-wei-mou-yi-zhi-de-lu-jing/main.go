package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
	. "github.com/9d77v/leetcode/pkg/algorithm/stack"
)

/*
题目：二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
*/

/*
运行时间：0 ms	内存消耗：2 MB
*/
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	var stack Stack = NewSliceStack(1)
	var findPathSum func(root *TreeNode, sum int, stack Stack, curSum int)
	findPathSum = func(root *TreeNode, sum int, stack Stack, curSum int) {
		curSum += root.Val
		stack.Push(root.Val)
		if curSum == sum && isLeft(root) {
			arr := make([]int, stack.Len())
			stack.Iterator(func(i int, v interface{}) {
				arr[i] = v.(int)
			})
			result = append(result, arr)
		}
		if root.Left != nil {
			findPathSum(root.Left, sum, stack, curSum)
		}
		if root.Right != nil {
			findPathSum(root.Right, sum, stack, curSum)
		}
		stack.Pop()
	}
	findPathSum(root, sum, stack, 0)
	return result
}

func isLeft(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
