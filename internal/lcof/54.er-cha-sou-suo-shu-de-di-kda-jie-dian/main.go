package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
)

/*
题目：二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第k大的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
*/

/*
方法一：逆中序遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：6.3 MB
*/
func kthLargest(root *TreeNode, k int) (result int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			result = root.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return
}
