package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
)

/*
题目：对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

限制：
0 <= 节点个数 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
*/

func isSymmetric(root *TreeNode) bool {
	return isSymmetricCore(root, root)
}

func isSymmetricCore(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetricCore(root1.Left, root2.Right) && isSymmetricCore(root2.Left, root1.Right)
}
