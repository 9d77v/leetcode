package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
)

/*
题目：树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

限制：

0 <= 节点个数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
*/

func isSubStructure(A *TreeNode, B *TreeNode) (result bool) {
	if A != nil && B != nil {
		if A.Val == B.Val {
			result = doesTreeAHasTreeB(A, B)
		}
		if !result {
			result = isSubStructure(A.Left, B)
		}
		if !result {
			result = isSubStructure(A.Right, B)
		}
	}
	return
}

func doesTreeAHasTreeB(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return doesTreeAHasTreeB(A.Left, B.Left) && doesTreeAHasTreeB(A.Right, B.Right)
}
