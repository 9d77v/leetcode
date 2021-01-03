package singlylinkedlist

//ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//NewList ..
func NewList(data []int) *ListNode {
	n := &ListNode{}
	m := n
	for i, num := range data {
		n.Val = num
		if i != len(data)-1 {
			n.Next = &ListNode{}
			n = n.Next
		}
	}
	return m
}

//ToArray 转数组
func (l *ListNode) ToArray() []int {
	if l == nil {
		return []int{}
	}
	res := make([]int, 0)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
