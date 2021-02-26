package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：最长端流子数组
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度。

提示：

1 <= A.length <= 40000
0 <= A[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-turbulent-subarray
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：76 ms	内存消耗：7.2 MB
*/
func maxTurbulenceSize(arr []int) int {
	left, max := 0, 1
	for right := 0; right < len(arr)-1; right++ {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
		} else {
			if !((arr[right] > arr[right-1] && arr[right] > arr[right+1]) || (arr[right] < arr[right-1] && arr[right] < arr[right+1])) {
				left = right
			}
		}
		max = Max(max, right-left+2)
	}
	return max
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：80 ms	内存消耗：7.4 MB
*/
func maxTurbulenceSizeFunc2(arr []int) int {
	n := len(arr)
	up, down, max := make([]int, n), make([]int, n), 1
	for i := 0; i < n; i++ {
		up[i] = 1
		down[i] = 1
	}
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			up[i] = down[i-1] + 1
		} else if arr[i-1] > arr[i] {
			down[i] = up[i-1] + 1
		}
		max = MaxArr(max, up[i], down[i])
	}
	return max
}

/*
方法二：动态规划优化
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：80 ms	内存消耗：7 MB
*/
func maxTurbulenceSizeFunc3(arr []int) int {
	up, down, max := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			up, down = down+1, 1
		} else if arr[i-1] > arr[i] {
			down, up = up+1, 1
		} else {
			up, down = 1, 1
		}
		max = MaxArr(max, up, down)
	}
	return max
}
