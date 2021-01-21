package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。

如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。

给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。

提示：

1 <= stones.length <= 1000
0 <= xi, yi <= 104
不会有两块石头放在同一个坐标点上

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column
*/

/*
方法一：并查集(数组)
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：20 ms	内存消耗：7.4 MB
*/
func removeStonesFunc1(stones [][]int) (res int) {
	var uf UnionFind = NewArrayUnionFind(20000)
	for _, stone := range stones {
		uf.Union(stone[0]+10000, stone[1])
	}
	return len(stones) - uf.Count()
}

/*
方法一：并查集(map)
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：20 ms	内存消耗：7.4 MB
*/
func removeStonesFunc2(stones [][]int) (res int) {
	var uf UnionFind = NewMapUnionFind(0)
	for _, stone := range stones {
		uf.Union(stone[0]+10000, stone[1])
	}
	return len(stones) - uf.Count()
}
