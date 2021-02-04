package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给你一个字符串 s，以及该字符串中的一些「索引对」数组 pairs，其中 pairs[i] = [a, b] 表示字符串中的两个索引（编号从 0 开始）。

你可以 任意多次交换 在 pairs 中任意一对索引处的字符。

返回在经过若干次交换后，s 可以变成的按字典序最小的字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-string-with-swaps
*/

/*
方法一：并查集
时间复杂度：О(n²㏒n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func smallestStringWithSwaps(s string, pairs [][]int) string {
	if s == "" {
		return s
	}
	n := len(s)
	var uf UnionFind = NewArrayUnionFind(n)
	for _, v := range pairs {
		uf.Union(v[0], v[1])
	}

	arr := []byte(s)
	hashMap := make(map[int]Heap, n)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		if hashMap[root] != nil {
			hashMap[root].PushItem(arr[i])
		} else {
			minHeap := NewMinHeap()
			minHeap.PushItem(arr[i])
			hashMap[root] = minHeap
		}
	}
	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		result = append(result, hashMap[root].PopItem().(byte))
	}
	return string(result)
}
