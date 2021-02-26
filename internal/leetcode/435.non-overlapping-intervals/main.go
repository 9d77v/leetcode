package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：无重叠区间
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-overlapping-intervals/
*/

/*
方法一：动态规划
时间复杂度：О(n²)
空间复杂度：О(n)
运行时间：276 ms	内存消耗：4.1 MB
*/
func eraseOverlapIntervalsFunc1(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	return n - MaxArr(dp...)
}

/*
方法二：贪心算法
时间复杂度：О(nlogn)
空间复杂度：О(logn)
运行时间：8 ms	内存消耗：3.9 MB
*/
func eraseOverlapIntervalsFunc2(intervals [][]int) int {
	return len(intervals) - OverlapIntervalSchedule(intervals)
}
