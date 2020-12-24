package main

/*
题目：给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明：
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

提示：
-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
*/

/*
方法一：双指针 / 从后往前
时间复杂度：О(n+m)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.3 MB
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	p, p1, p2 := m+n-1, m-1, n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
