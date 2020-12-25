package main

import (
	"sort"
)

/*
题目：
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

提示：
1 <= g.length <= 3 * 10^4
0 <= s.length <= 3 * 10^4
1 <= g[i], s[j] <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/assign-cookies
*/

/*
方法一：排序+贪心算法
时间复杂度： О(nlogn)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：6.2 MB
*/
func findContentChildrenFunc1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j, n, m := 0, 0, len(g), len(s)
	for i < n && j < m {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
