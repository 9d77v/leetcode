package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
	. "github.com/9d77v/leetcode/pkg/algorithm/queue"
)

/*
题目：
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
*/

func levelOrder(root *TreeNode) [][]int {
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
	return result
}
