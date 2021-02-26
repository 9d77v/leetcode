package main

/*
题目：较大分组的位置
在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。

例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。

分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。上例中的 "xxxx" 分组用区间表示为 [3,6] 。

我们称所有包含大于或等于三个连续字符的分组为 较大分组 。

找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。

提示：
1 <= s.length <= 1000
s 仅含小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/positions-of-large-groups
*/

/*
方法一：一次遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.6 MB
*/
func largeGroupPositions(s string) [][]int {
	if len(s) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	ch, start := s[0], 0
	for i := 1; i < len(s); i++ {
		if s[i] != ch {
			if i-start >= 3 {
				res = append(res, []int{start, i - 1})
			}
			start, ch = i, s[i]
		} else if len(s)-1 == i && i-start >= 2 {
			res = append(res, []int{start, i})
		}
	}
	return res
}
