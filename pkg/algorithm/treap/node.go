package treap

import "math/rand"

//Node 数堆节点
type Node struct {
	priority int
	key      int
	size     int
	children [2]*Node
}

func (n *Node) cmp(b int) int {
	switch {
	case b < n.key:
		return 0
	case b > n.key:
		return 1
	default:
		return -1
	}
}

func (n *Node) rotate(d int) *Node {
	x := n.children[d^1]
	n.children[d^1] = x.children[d]
	x.children[d] = n
	return x
}

//Treap 数堆
type Treap struct {
	root *Node
}

func (t *Treap) ins(n *Node, key int) *Node {
	if n == nil {
		return &Node{priority: rand.Int(), key: key, size: 1}
	}
	if d := n.cmp(key); d >= 0 {
		n.children[d] = t.ins(n.children[d], key)
		if n.children[d].priority > n.priority {
			n = n.rotate(d ^ 1)
		}
	} else {
		n.size++
	}
	return n
}

func (t *Treap) del(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if d := node.cmp(key); d >= 0 {
		node.children[d] = t.del(node.children[d], key)
	} else {
		if node.size > 1 {
			node.size--
		} else {
			if node.children[1] == nil {
				return node.children[0]
			}
			if node.children[0] == nil {
				return node.children[1]
			}
			d = 0
			if node.children[0].priority > node.children[1].priority {
				d = 1
			}
			node = node.rotate(d)
			node.children[d] = t.del(node.children[d], key)
		}
	}
	return node
}

//Insert 新增
func (t *Treap) Insert(key int) {
	t.root = t.ins(t.root, key)
}

//Delete 删除
func (t *Treap) Delete(key int) {
	t.root = t.del(t.root, key)
}

//Min 最小值
func (t *Treap) Min() (min int) {
	for node := t.root; node != nil; node = node.children[0] {
		min = node.key
	}
	return
}

//Max 最大值
func (t *Treap) Max() (max int) {
	for node := t.root; node != nil; node = node.children[1] {
		max = node.key
	}
	return
}
