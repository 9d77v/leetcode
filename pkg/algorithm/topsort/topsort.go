package topsort

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/queue"
)

//BfsTopSort 拓扑排序广度优先搜索
func BfsTopSort(n int, prerequisites [][]int, fn func(u int)) {
	graph := make([][]int, n)
	indegrees := make([]int, n)
	for _, info := range prerequisites {
		graph[info[1]] = append(graph[info[1]], info[0])
		indegrees[info[0]]++
	}
	var queue Queue = NewSliceQueue(n)
	for i, indegree := range indegrees {
		if indegree == 0 {
			queue.Push(i)
		}
	}
	queue.BFS(func(u interface{}) {
		if fn != nil {
			fn(u.(int))
		}
		for _, v := range graph[u.(int)] {
			indegrees[v]--
			if indegrees[v] == 0 {
				queue.Push(v)
			}
		}
	})
}

//DfsTopSort 拓扑排序广度优先搜索
func DfsTopSort(n int, prerequisites [][]int, valid bool, fn func(u int)) {
	graph := make([][]int, n)
	visited := make([]int, n)
	for _, info := range prerequisites {
		graph[info[1]] = append(graph[info[1]], info[0])
	}
	var dfs func(u int, fn func(u int))
	dfs = func(u int, fn func(u int)) {
		visited[u] = 1
		for _, v := range graph[u] {
			if visited[v] == 0 {
				dfs(v, fn)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		if fn != nil {
			fn(u)
		}
	}
	for i := 0; i < len(graph) && valid; i++ {
		if visited[i] == 0 {
			dfs(i, fn)
		}
	}
}
