package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：最长递增子序列

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

 提示：

1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法一：动态规划,以nums[i]结尾的上升子序列
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：120 ms	内存消耗：3.7 MB
*/
func lengthOfLIS(nums []int) (max int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = MaxArr(dp[i], dp[j]+1)
			}
		}
	}
	return MaxArr(dp...)
}

/*
方法二：贪心+二分查找,长度为i+1的所有上升子序列的结尾最小值
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：8 ms	内存消耗：3.5 MB
*/
func lengthOfLISFunc2(nums []int) (max int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	d, end := make([]int, n), 0
	d[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[end] {
			end++
			d[end] = nums[i]
		} else {
			left := sort.SearchInts(d[:end+1], nums[i])
			d[left] = nums[i]
		}
	}
	end++
	return end
}
