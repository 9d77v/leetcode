package singlylinkedlist

//ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//NewListNode ..
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}
