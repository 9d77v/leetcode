package algorithm

//BfsTopSort 拓扑排序广度优先搜索
func BfsTopSort(graph [][]int, indegrees []int, fn func(u int)) {
	queue := []int{}
	for i, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		if fn != nil {
			fn(u)
		}
		for _, v := range graph[u] {
			indegrees[v]--
			if indegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
}

//DfsTopSort 拓扑排序广度优先搜索
func DfsTopSort(graph [][]int, visited []int, valid bool, fn func(u int)) {
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
