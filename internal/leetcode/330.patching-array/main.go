package main

/*
题目：按要求补齐数组
给定一个已排序的正整数数组 nums，和一个正整数 n 。从 [1, n] 区间内选取任意个数字补充到 nums 中，使得 [1, n] 区间内的任何数字都可以用 nums 中某几个数字的和来表示。请输出满足上述要求的最少需要补充的数字个数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/patching-array
*/

/*
方法一：贪心算法
时间复杂度：О(m+㏒n)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：3.3 MB

思路：
最少数字全覆盖
1,2  1->3
1,2,4 1->7
1,2,4,8 1->15
*/
func minPatchesFunc1(nums []int, n int) (res int) {
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && x >= nums[i] {
			x += nums[i]
			i++
		} else {
			x *= 2
			res++
		}
	}
	return
}
