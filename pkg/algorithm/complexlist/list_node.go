package complexlist

//Node 复杂链表
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//NewComplexList 初始化复杂链表
func NewComplexList(arr [][]*int) *Node {
	n := len(arr)
	if n == 0 {
		return nil
	}
	nodes := make([]*Node, n)
	for i, v := range arr {
		nodes[i] = &Node{
			Val: *v[0],
		}
	}
	for i := 0; i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	for i, v := range arr {
		if v[1] != nil {
			nodes[i].Random = nodes[*v[1]]
		}
	}
	return nodes[0]
}
