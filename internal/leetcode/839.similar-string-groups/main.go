package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：相似字符串组
如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。

例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给你一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个字母异位词。请问 strs 中有多少个相似字符串组？

提示：

1 <= strs.length <= 100
1 <= strs[i].length <= 1000
sum(strs[i].length) <= 2 * 104
strs[i] 只包含小写字母。
strs 中的所有单词都具有相同的长度，且是彼此的字母异位词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/similar-string-groups
*/

/*
方法一：并查集
时间复杂度：О(n^2m+nlogn)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.6 MB
*/
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
