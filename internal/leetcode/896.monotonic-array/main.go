package main

import "sort"

/*
题目：
如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A 是单调数组时返回 true，否则返回 false。

提示：

1 <= A.length <= 50000
-100000 <= A[i] <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/monotonic-array
*/

/*
方法一：两次遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：72 ms	内存消耗：7.8 MB
*/
func isMonotonic(A []int) bool {
	n := len(A)
	if n <= 2 {
		return true
	}
	flag := true
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			flag = false
			break
		}
	}
	if flag {
		return true
	}
	flag = true
	for i := 1; i < n; i++ {
		if A[i] < A[i-1] {
			flag = false
			break
		}
	}
	return flag
}

/*
方法二：两次遍历,使用系统方法,修改数组
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：76 ms	内存消耗：7.9 MB
*/
func isMonotonicFunc2(A []int) bool {
	return sort.IntsAreSorted(A) || sort.IsSorted(sort.Reverse(sort.IntSlice(A)))
}

/*
方法三：一次遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：76 ms	内存消耗：8.1 MB
*/
func isMonotonicFunc3(A []int) bool {
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			inc = false
		}
		if A[i] > A[i-1] {
			dec = false
		}
	}
	return inc || dec
}

/*
方法四：一次遍历,头尾确定递增或递减
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：76 ms	内存消耗：7.8 MB
*/
func isMonotonicFunc4(A []int) bool {
	n := len(A)
	if A[0] < A[n-1] {
		for i := 1; i < n; i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
	} else {
		for i := 1; i < n; i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
	}
	return true
}
