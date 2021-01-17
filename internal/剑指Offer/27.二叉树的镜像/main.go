package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
)

/*
题目：
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

限制：
0 <= 节点个数 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
*/

func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		mirrorTree(root.Left)
		mirrorTree(root.Right)
	}
	return root
}
