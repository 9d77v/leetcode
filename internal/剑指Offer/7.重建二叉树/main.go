package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/binarytree"
	. "github.com/9d77v/leetcode/pkg/algorithm/stack"
)

/*
题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
限制：
0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
*/

/*
方法一：递归
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：4.2 MB
*/
func buildTreeFunc1(preorder []int, inorder []int) *TreeNode {
	preSize, inSize := len(preorder), len(inorder)
	if preSize == 0 || inSize == 0 || preSize != inSize {
		return nil
	}
	rootIndex := getRootIndex(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTreeFunc1(preorder[1:rootIndex+1], inorder[:rootIndex]),
		Right: buildTreeFunc1(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}

func getRootIndex(inorder []int, val int) int {
	rootIndex := 0
	for i, v := range inorder {
		if v == val {
			rootIndex = i
			break
		}
	}
	return rootIndex
}

/*
方法二：栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：3.5 MB
*/
func buildTreeFunc2(preorder []int, inorder []int) *TreeNode {
	preSize, inSize := len(preorder), len(inorder)
	if preSize == 0 || inSize == 0 || preSize != inSize {
		return nil
	}
	var stack Stack = NewSliceStack(preSize)
	root := NewTreeNode(preorder[0])
	cur := root
	for i, j := 1, 0; i < preSize; i++ {
		if cur.Val != inorder[j] {
			cur.Left = NewTreeNode(preorder[i])
			stack.Push(cur)
			cur = cur.Left
		} else {
			j++
			for !stack.Empty() && stack.Peek().(*TreeNode).Val == inorder[j] {
				cur = stack.Pop().(*TreeNode)
				j++
			}
			cur.Right = NewTreeNode(preorder[i])
			cur = cur.Right
		}
	}
	return root
}
