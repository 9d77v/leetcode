package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
)

/*
题目：平衡二叉树
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

提示：

节点总数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
*/

/*
方法一：调用二叉树最大深度，重复计算
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：6.1 MB
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := maxDepth(root.Left) - maxDepth(root.Right)
	if diff > 1 || diff < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/*
方法二：递归
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：6.1 MB
*/
func isBalancedFunc2(root *TreeNode) bool {
	depth := 0
	return isBalancedCore(root, &depth)
}

func isBalancedCore(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	left, right := 0, 0
	if isBalancedCore(root.Left, &left) && isBalancedCore(root.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			*depth = 1 + Max(left, right)
			return true
		}
	}
	return false
}
