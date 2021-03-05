package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/russian-doll-envelopes
*/

/*
方法一：动态规划
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：5.3 MB
*/
func maxEnvelopes(envelopes [][]int) (max int) {
	n := len(envelopes)
	if n <= 1 {
		return n
	}
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return (a[0] < b[0]) || (a[0] == b[0] && a[1] > b[1])
	})
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = MaxArr(dp[i], dp[j]+1)
			}
		}
	}
	return MaxArr(dp...)
}

/*
方法二：贪心+二分
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：28 ms	内存消耗：6.5 MB
*/
func maxEnvelopesFunc2(envelopes [][]int) (max int) {
	n := len(envelopes)
	if n <= 1 {
		return n
	}
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return (a[0] < b[0]) || (a[0] == b[0] && a[1] > b[1])
	})
	d, end := make([]int, n), 0
	d[end] = envelopes[0][1]
	for i := 1; i < n; i++ {
		if envelopes[i][1] > d[end] {
			end++
			d[end] = envelopes[i][1]
		} else {
			left := sort.SearchInts(d[:end+1], envelopes[i][1])
			d[left] = envelopes[i][1]
		}
	}
	end++
	return end
}
