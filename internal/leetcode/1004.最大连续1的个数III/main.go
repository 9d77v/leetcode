package main

/*
题目：
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

提示：

1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii/
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：68 ms	内存消耗：6.9 MB
*/
func longestOnes(A []int, K int) int {
	left, zeroCnt := 0, 0
	for _, num := range A {
		if num == 0 {
			zeroCnt++
		}
		if zeroCnt > K {
			if A[left] == 0 {
				zeroCnt--
			}
			left++
		}
	}
	return len(A) - left
}
