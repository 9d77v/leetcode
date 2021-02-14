package main

/*
题目：
给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。

（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）

返回 A 中好子数组的数目。

提示：

1 <= A.length <= 20000
1 <= A[i] <= A.length
1 <= K <= A.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarrays-with-k-different-integers
*/

/*
方法一：双指针
时间复杂度：О(n+ClogC)
空间复杂度：О(C)
运行时间：52 ms	内存消耗：6.9 MB
*/
func subarraysWithKDistinct(A []int, K int) int {
	return atMostKDistinct(A, K) - atMostKDistinct(A, K-1)
}

func atMostKDistinct(A []int, K int) int {
	left, count, res, freq := 0, 0, 0, make([]int, len(A)+1)
	for right := range A {
		if freq[A[right]] == 0 {
			count++
		}
		freq[A[right]]++
		for count > K {
			freq[A[left]]--
			if freq[A[left]] == 0 {
				count--
			}
			left++
		}
		res += right - left + 1
	}
	return res
}
