package main

import (
	"math"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。

提示：

1 <= points.length <= 1000
-106 <= xi, yi <= 106
所有点 (xi, yi) 两两不同。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-to-connect-all-points
*/

/*
方法一：Kruskal 算法
时间复杂度：О(n²log(n))
空间复杂度：О(n²)
运行时间：392 ms	内存消耗：27 MB
*/
func minCostConnectPointsFunc1(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}
	edges := getEdges(points)
	sortByLen(edges)
	return calcCost(edges, n)
}

func getEdges(points [][]int) [][3]int {
	edges := make([][3]int, 0, len(points)*(len(points)-1))
	n := len(points)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, [3]int{i, j, dist(points[i], points[j])})
		}
	}
	return edges
}

func dist(x, y []int) int {
	return Abs(x[0]-y[0]) + Abs(x[1]-y[1])
}

func sortByLen(edges [][3]int) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
}

func calcCost(edges [][3]int, n int) (cost int) {
	var uf UnionFind = NewArrayUnionFindWithRank(n, RankSize)
	left := n - 1
	for _, edge := range edges {
		if uf.Union(edge[0], edge[1]) {
			cost += edge[2]
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}

/*
方法二：建图优化的Kruskal 算法
时间复杂度：О(nlog(n))
空间复杂度：О(n)
运行时间：16 ms	内存消耗：6.5 MB
*/
func minCostConnectPointsFunc2(points [][]int) int {
	n := len(points)
	edges := getEdges2(points, n)
	sortByLen(edges)
	return calcCost(edges, n)
}

func getEdges2(points [][]int, n int) [][3]int {
	for i, p := range points {
		points[i] = append(p, i)
	}
	edges := [][3]int{}
	build := func() {
		sort.Slice(points, func(i, j int) bool { a, b := points[i], points[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

		// 离散化 y-x
		type pair struct{ v, i int }
		ps := make([]pair, n)
		for i, p := range points {
			ps[i] = pair{p[1] - p[0], i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
		kth := make([]int, n)
		k := 1
		kth[ps[0].i] = k
		for i := 1; i < n; i++ {
			if ps[i].v != ps[i-1].v {
				k++
			}
			kth[ps[i].i] = k
		}

		t := newFenwickTree(k + 1)
		for i := n - 1; i >= 0; i-- {
			p := points[i]
			pos := kth[i]
			if j := t.query(pos); j != -1 {
				q := points[j]
				edges = append(edges, [3]int{p[2], q[2], dist(p, q)})
			}
			t.update(pos, p[0]+p[1], i)
		}
	}

	build()
	for _, p := range points {
		p[0], p[1] = p[1], p[0]
	}
	build()
	for _, p := range points {
		p[0] = -p[0]
	}
	build()
	for _, p := range points {
		p[0], p[1] = p[1], p[0]
	}
	build()
	return edges
}

type fenwickTree struct {
	tree, idRec []int
}

func newFenwickTree(n int) *fenwickTree {
	tree := make([]int, n)
	idRec := make([]int, n)
	for i := range tree {
		tree[i], idRec[i] = math.MaxInt64, -1
	}
	return &fenwickTree{tree, idRec}
}

func (f *fenwickTree) update(pos, val, id int) {
	for ; pos > 0; pos &= pos - 1 {
		if val < f.tree[pos] {
			f.tree[pos], f.idRec[pos] = val, id
		}
	}
}

func (f *fenwickTree) query(pos int) int {
	minVal, minID := math.MaxInt64, -1
	for ; pos < len(f.tree); pos += pos & -pos {
		if f.tree[pos] < minVal {
			minVal, minID = f.tree[pos], f.idRec[pos]
		}
	}
	return minID
}
