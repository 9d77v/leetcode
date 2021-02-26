package main

/*
题目：K连续位的最小翻转次数
在仅包含 0 和 1 的数组 A 中，一次 K 位翻转包括选择一个长度为 K 的（连续）子数组，同时将子数组中的每个 0 更改为 1，而每个 1 更改为 0。

返回所需的 K 位翻转的最小次数，以便数组没有值为 0 的元素。如果不可能，返回 -1。

提示：

1 <= A.length <= 30000
1 <= K <= A.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips
*/

/*
方法一：差分数组
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：112 ms	内存消耗：7 MB
*/
func minKBitFlips(A []int, K int) (result int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt += diff[i]
		if (v+revCnt)&1 == 0 {
			if i+K > n {
				return -1
			}
			result++
			revCnt++
			diff[i+K]--
		}
	}
	return
}

/*
方法二：差分数组
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：108 ms	内存消耗：7 MB
*/
func minKBitFlipsFunc2(A []int, K int) (result int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v == revCnt {
			if i+K > n {
				return -1
			}
			result++
			revCnt ^= 1
			diff[i+K] ^= 1
		}
	}
	return
}

/*
方法三：滑动窗口
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：140/108 ms	内存消耗：7.3 MB
*/
func minKBitFlipsFunc3(A []int, K int) (result int) {
	n := len(A)
	revCnt := 0
	for i, v := range A {
		if i >= K && A[i-K] > 1 {
			revCnt ^= 1
			A[i-K] -= 2 //不要求还原，不用执行
		}
		if v == revCnt {
			if i+K > n {
				return -1
			}
			result++
			revCnt ^= 1
			A[i] += 2
		}
	}
	return
}
