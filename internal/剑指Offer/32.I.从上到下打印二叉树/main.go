package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
	. "github.com/9d77v/leetcode/pkg/algorithm/queue"
)

/*
题目：
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
*/
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	var queue Queue = NewSliceQueue(1, root)
	queue.BFS(func(front interface{}) {
		node := front.(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
	})
	return result
}
