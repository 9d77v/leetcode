package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

func numSimilarGroups(strs []string) int {
	n := len(strs)
	var uf UnionFind = NewArrayUnionFind(n, RankSize)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if !uf.IsConnected(i, j) && isSimiliar(strs[i], strs[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count()
}

func isSimiliar(a, b string) bool {
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
}
