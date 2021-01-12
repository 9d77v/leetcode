package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

提示：
1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj 由小写英文字母与数字组成
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/evaluate-division
*/

/*
方法一：广度优先搜索
时间复杂度：О(ML+Q⋅(L+M))
空间复杂度：О(NL+M)
运行时间：0 ms	内存消耗：2.1 MB
*/
func calcEquationFunc1(equations [][]string, values []float64, queries [][]string) []float64 {
	idMap := buildIDMap(equations)
	graph := buildEdgeGraph(idMap, equations, values)
	res := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := idMap[q[0]]
		end, hasE := idMap[q[1]]
		if !hasS || !hasE {
			res[i] = -1
		} else {
			res[i] = bfs(graph, start, end)
		}
	}
	return res
}

func buildIDMap(equations [][]string) map[string]int {
	idMap := make(map[string]int, 2*len(equations))
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := idMap[a]; !has {
			idMap[a] = len(idMap)
		}
		if _, has := idMap[b]; !has {
			idMap[b] = len(idMap)
		}
	}
	return idMap
}

func buildEdgeGraph(idMap map[string]int, equations [][]string, values []float64) [][]edge {
	graph := make([][]edge, len(idMap))
	for i, eq := range equations {
		v, w := idMap[eq[0]], idMap[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v, 1 / values[i]})
	}
	return graph
}

type edge struct {
	to     int
	weight float64
}

func bfs(graph [][]edge, start, end int) float64 {
	ratios := make([]float64, len(graph))
	ratios[start] = 1
	queue := []int{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == end {
			return ratios[v]
		}
		for _, e := range graph[v] {
			if w := e.to; ratios[w] == 0 {
				ratios[w] = ratios[v] * e.weight
				queue = append(queue, w)
			}
		}
	}
	return -1
}

/*
方法二：flyod算法
时间复杂度：О(ML+N^3+QL)
空间复杂度：О(NL+N^2)
运行时间：0 ms	内存消耗：2.1 MB
*/
func calcEquationFunc2(equations [][]string, values []float64, queries [][]string) []float64 {
	idMap := buildIDMap(equations)
	graph := buildGraph(idMap, equations, values)
	flyod(graph)
	res := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := idMap[q[0]]
		end, hasE := idMap[q[1]]
		if !hasS || !hasE || graph[start][end] == 0 {
			res[i] = -1
		} else {
			res[i] = graph[start][end]
		}
	}
	return res
}

func buildGraph(idMap map[string]int, equations [][]string, values []float64) [][]float64 {
	graph := make([][]float64, len(idMap))
	for i := range graph {
		graph[i] = make([]float64, len(idMap))
	}
	for i, eq := range equations {
		v, w := idMap[eq[0]], idMap[eq[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1 / values[i]
	}
	return graph
}

func flyod(graph [][]float64) {
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] > 0 && graph[k][j] > 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}
}

/*
方法三：并查集
时间复杂度：О(ML+N+MlogN+Q⋅(L+logN))
空间复杂度：О(NL)
运行时间：0 ms	内存消耗：2.1 MB
*/
func calcEquationFunc3(equations [][]string, values []float64, queries [][]string) []float64 {
	equationsSize := len(equations)
	unionFind := NewWeightUnionFind(2 * equationsSize)
	idMap := make(map[string]int, 2*equationsSize)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if _, has := idMap[a]; !has {
			idMap[a] = len(idMap)
		}
		if _, has := idMap[b]; !has {
			idMap[b] = len(idMap)
		}
		unionFind.Union(idMap[a], idMap[b], values[i])
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		id1, hasS := idMap[q[0]]
		id2, hasE := idMap[q[1]]
		if !hasS || !hasE {
			res[i] = -1
		} else {
			res[i] = unionFind.IsConnected(id1, id2)
		}
	}
	return res
}
