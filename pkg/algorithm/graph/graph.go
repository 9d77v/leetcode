package graph

import (
	"fmt"
	"strings"
)

//Graph 无向图
type Graph struct {
	adj [][]int //邻接表
	v   int     //顶点数目
	e   int     //边的数目
}

//NewGraph 初始化无向图
func NewGraph(v int) *Graph {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}
	return &Graph{
		adj: adj,
		v:   v,
	}
}

//NewGrapWithEdges 初始化存在边的无向图
func NewGrapWithEdges(v int, edges [][]int) *Graph {
	g := NewGraph(v)
	g.e = len(edges)
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}
	return g
}

//AddEdge 添加一条边v-w
func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

//V 顶点数
func (g *Graph) V() int {
	return g.v
}

//E 边数
func (g *Graph) E() int {
	return g.e
}

//Adj 邻接表
func (g *Graph) Adj() [][]int {
	return g.adj
}

//Degree 获取v的度数
func (g *Graph) Degree(v int) int {
	return len(g.adj[v])
}

//MaxDegree 获取所有顶点的最大度数
func (g *Graph) MaxDegree() (count int) {
	for _, v := range g.adj {
		if count < len(v) {
			count = len(v)
		}
	}
	return
}

//AvgDegree 获取所有顶点的平均度数
func (g *Graph) AvgDegree() int {
	return 2 * g.e / g.v
}

//NumOfSelfLoops 获取自环的个数
func (g *Graph) NumOfSelfLoops() (count int) {
	for i, v := range g.adj {
		for _, w := range v {
			if i == w {
				count++
			}
		}
	}
	return count / 2
}

//ToString 字符串表示
func (g *Graph) ToString() string {
	builder := new(strings.Builder)
	builder.WriteString(fmt.Sprintf("vertices: %d, edges: %d\n", g.v, g.e))
	for i, v := range g.adj {
		builder.WriteString(fmt.Sprintf("%d: ", i))
		for _, w := range v {
			builder.WriteString(fmt.Sprintf("%d ", w))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
