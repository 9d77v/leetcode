package main

import "sort"

/*
题目：
给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。

每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。

请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。

提示：

1 <= nums1.length, nums2.length <= 105
1 <= nums1[i], nums2[i] <= 6
*/

func minOperations(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	min := Min(m, n)
	if (min == m && m*6 < n) || (min == n && n*6 < m) {
		return -1
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	sum1, sum2 := 0, 0
	for _, num := range nums1 {
		sum1 += num
	}
	for _, num := range nums2 {
		sum2 += num
	}
	return 0
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
