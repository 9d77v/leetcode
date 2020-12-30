package binarytree

import "fmt"

//TreeNode  二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//NewTreeNode 新节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

//PreorderTraverse 前序遍历 根->左->右
func (m *TreeNode) PreorderTraverse() {
	if m == nil {
		return
	}
	fmt.Print(m.Val, " ")
	m.Left.PreorderTraverse()
	m.Right.PreorderTraverse()
	return
}

//InorderTraverse 中序遍历 左->根->右
func (m *TreeNode) InorderTraverse() {
	if m == nil {
		return
	}
	m.Left.InorderTraverse()
	fmt.Print(m.Val, " ")
	m.Right.InorderTraverse()
	return
}

//PostTraverse 中序遍历 左->右->根
func (m *TreeNode) PostTraverse() {
	if m == nil {
		return
	}
	m.Left.InorderTraverse()
	m.Right.InorderTraverse()
	fmt.Print(m.Val, " ")
	return
}

//BFS 广度优先遍历
func (m *TreeNode) BFS() []interface{} {
	var result []interface{}
	var nodes []TreeNode = []TreeNode{*m}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return result
}

//BFSWithNil 广度优先遍历 带nil
func (m *TreeNode) BFSWithNil() []interface{} {
	var result []interface{}
	var nodes []*TreeNode = []*TreeNode{m}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node == nil {
			if notAllNil(nodes) {
				result = append(result, nil)
			}
			continue
		}
		result = append(result, node.Val)
		nodes = append(nodes, node.Left)
		nodes = append(nodes, node.Right)
	}
	return result
}

func notAllNil(nodes []*TreeNode) bool {
	flag := false
	for _, v := range nodes {
		if v != nil {
			flag = true
			break
		}
	}
	return flag
}
