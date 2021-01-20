package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected
*/

/*
方法一：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：72 ms	内存消耗 10.2 MB
*/
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	var uf UnionFind = NewArrayUnionFindWithRank(n, RankSize)
	for _, conn := range connections {
		if uf.Union(conn[0], conn[1]) {
			n--
		}
	}
	return n - 1
}
