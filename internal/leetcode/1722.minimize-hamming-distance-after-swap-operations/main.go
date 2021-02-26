package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：执行交换操作后的最小汉明距离

给你两个整数数组 source 和 target ，长度都是 n 。还有一个数组 allowedSwaps ，其中每个 allowedSwaps[i] = [ai, bi] 表示你可以交换数组 source 中下标为 ai 和 bi（下标从 0 开始）的两个元素。注意，你可以按 任意 顺序 多次 交换一对特定下标指向的元素。

相同长度的两个数组 source 和 target 间的 汉明距离 是元素不同的下标数量。形式上，其值等于满足 source[i] != target[i] （下标从 0 开始）的下标 i（0 <= i <= n-1）的数量。

在对数组 source 执行 任意 数量的交换操作后，返回 source 和 target 间的 最小汉明距离 。

提示：
n == source.length == target.length
1 <= n <= 105
1 <= source[i], target[i] <= 105
0 <= allowedSwaps.length <= 105
allowedSwaps[i].length == 2
0 <= ai, bi <= n - 1
ai != bi

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimize-hamming-distance-after-swap-operations
*/

/*
方法一：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：296 ms	内存消耗 27.7 MB
*/
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) (result int) {
	var uf UnionFind = NewMapUnionFind(0)
	for _, v := range allowedSwaps {
		uf.Union(v[0], v[1])
	}

	type tmp struct {
		iArr []int
		vMap map[int]int
	}

	rootMap := make(map[int]*tmp, 0)
	for i, v := range target {
		if source[i] == v {
			result++
			continue
		}
		if uf.Has(i) {
			rootID := uf.Find(i)
			if _, ok := rootMap[rootID]; !ok {
				iArr := make([]int, 1, uf.Rank(rootID))
				iArr[0] = i
				rootMap[rootID] = &tmp{iArr, map[int]int{v: 1}}
			} else {
				rootMap[rootID].iArr = append(rootMap[rootID].iArr, i)
				rootMap[rootID].vMap[v]++
			}
		}
	}
	for rootID, v := range rootMap {
		for _, i := range v.iArr {
			if rootMap[rootID].vMap[source[i]] > 0 {
				rootMap[rootID].vMap[source[i]]--
				result++
			}
		}
	}
	return len(target) - result
}
