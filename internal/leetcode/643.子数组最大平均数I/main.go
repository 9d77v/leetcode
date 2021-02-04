package main

import "math"

/*
题目：
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

提示：

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-average-subarray-i/
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：136ms	内存消耗： 7.3MB
*/
func findMaxAverage(nums []int, k int) float64 {
	preSum, max := 0, math.Inf(-1)
	for i, num := range nums {
		preSum += num
		if i >= k {
			preSum -= nums[i-k]
		}
		if i >= k-1 {
			max = math.Max(max, float64(preSum)/float64(k))
		}
	}
	return max
}
