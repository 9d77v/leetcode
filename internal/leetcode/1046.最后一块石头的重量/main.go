package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
题目：
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight
*/

/*
方法一：贪心
时间复杂度：О(n²㏒n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func lastStoneWeightFun1(stones []int) int {
	n := len(stones)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for stones[n-2] > 0 {
		stones[n-1], stones[n-2] = stones[n-1]-stones[n-2], 0
		sort.Ints(stones)
	}
	return stones[n-1]
}

/*
方法二：大顶堆
时间复杂度：О(n㏒n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func lastStoneWeightFun2(stones []int) int {
	n := len(stones)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return stones[0]
	}
	var hp Heap = NewMaxHeap()
	for _, v := range stones {
		hp.PushItem(v)
	}
	for hp.Len() > 1 {
		x, y := hp.PopItem().(int), hp.PopItem().(int)
		if x > y {
			hp.PushItem(x - y)
		}
	}
	if hp.Len() > 0 {
		return hp.PopItem().(int)
	}
	return 0
}
