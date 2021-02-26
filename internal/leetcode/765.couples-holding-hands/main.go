package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：情侣牵手
N 对情侣坐在连续排列的 2N 个座位上，想要牵到对方的手。 计算最少交换座位的次数，以便每对情侣可以并肩坐在一起。 一次交换可选择任意两人，让他们站起来交换座位。

人和座位用 0 到 2N-1 的整数表示，情侣们按顺序编号，第一对是 (0, 1)，第二对是 (2, 3)，以此类推，最后一对是 (2N-2, 2N-1)。

这些情侣的初始座位  row[i] 是由最初始坐在第 i 个座位上的人决定的。

说明:

len(row) 是偶数且数值在 [4, 60]范围内。
可以保证row 是序列 0...len(row)-1 的一个全排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/couples-holding-hands
*/

/*
方法一：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.1 MB
*/
func minSwapsCouples(row []int) int {
	n := len(row)
	var uf UnionFind = NewArrayUnionFind(n)
	for i := 0; i < n; i += 2 {
		if isCouple(row[i], row[i+1]) {
			uf.Union(row[i], row[i+1])
		} else {
			uf.Union(row[i], row[i])
			uf.Union(row[i+1], row[i+1])
		}
	}
	count := 0
	for uf.Count() != n>>1 {
		flag := false
		for i := 0; i < n-2; i += 2 {
			if !uf.IsConnected(row[i], row[i+1]) {
				for j := i + 2; j < n; j += 2 {
					if (isCouple(row[i], row[j]) && isCouple(row[i+1], row[j+1])) || (isCouple(row[i], row[j+1]) && isCouple(row[i+1], row[j])) {
						row[i+1], row[j] = row[j], row[i+1]
						uf.Union(row[i], row[i+1])
						uf.Union(row[j], row[j+1])
						count++
						flag = true
						break
					}
				}
			}
		}
		if !flag {
		checkCouple:
			for i := 0; i < n-2; i += 2 {
				if !uf.IsConnected(row[i], row[i+1]) {
					for j := i + 2; j < n; j += 2 {
						if isCouple(row[i], row[j]) {
							row[i+1], row[j] = row[j], row[i+1]
							uf.Union(row[i], row[i+1])
							count++
							break checkCouple
						}
						if isCouple(row[i], row[j+1]) {
							row[i+1], row[j+1] = row[j+1], row[i+1]
							uf.Union(row[i], row[i+1])
							count++
							break checkCouple
						}
					}
				}
			}
		}
	}
	return count
}

func isCouple(i, j int) bool {
	return (i&1 == 0 && j&1 == 1 && j-i == 1) || (i&1 == 1 && j&1 == 0 && i-j == 1)
}

/*
方法二：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func minSwapsCouplesFunc2(row []int) int {
	n := len(row)
	var uf UnionFind = NewArrayUnionFind(n / 2)
	for i := 0; i < n; i += 2 {
		uf.Union(row[i]>>1, row[i+1]>>1)
	}
	return n>>1 - uf.Count()
}
