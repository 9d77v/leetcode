package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
	. "github.com/9d77v/leetcode/pkg/algorithm/queue"
)

/*
题目：滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
*/

/*
方法零：循环遍历
时间复杂度：О(nk)
空间复杂度：O(n)
运行时间：超出时间限制
*/
func maxSlidingWindowFunc0(nums []int, k int) []int {
	tempArr, res, cur := make([]int, k), make([]int, 0, len(nums)-k+1), 0
	for i := 0; i < len(nums); i++ {
		tempArr[cur] = nums[i]
		if i >= k-1 {
			res = append(res, MaxArr(tempArr...))
		}
		cur = (cur + 1) % k
	}
	return res
}

/*
方法一：大顶堆/先序队列
时间复杂度：О(n㏒n)
空间复杂度：O(n)
运行时间：424 ms	内存消耗：16 MB
*/
func maxSlidingWindowFunc1(nums []int, k int) []int {
	n := len(nums)
	var hp Heap = NewMaxHeap(k)
	for i := 0; i < k; i++ {
		hp.PushItem([2]int{i, nums[i]})
	}
	res := make([]int, 1, n-k+1)
	res[0] = hp.Peek().([2]int)[1]
	for i := k; i < n; i++ {
		hp.PushItem([2]int{i, nums[i]})
		for hp.Peek().([2]int)[0] <= i-k {
			hp.PopItem()
		}
		res = append(res, hp.Peek().([2]int)[1])
	}
	return res
}

/*
方法二：单调队列
时间复杂度：О(n)
空间复杂度：O(k)
运行时间：284 ms	内存消耗：13 MB
*/
func maxSlidingWindowFunc2(nums []int, k int) []int {
	monotonicQueue := NewMonotonicQueue(NewSliceDeque(k), false)
	res := make([]int, 0, len(nums)-k+1)
	monotonicQueue.Execute(nums, func(fronIndex, frontValue int) {
		res = append(res, frontValue)
	})
	return res
}

/*
方法三：分块+预处理
时间复杂度：О(n)
空间复杂度：O(n)
运行时间：248 ms	内存消耗：10.1 MB
*/
func maxSlidingWindowFunc3(nums []int, k int) []int {
	n := len(nums)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)
	maxLeft[0] = nums[0]
	maxRight[n-1] = nums[n-1]
	for i, v := range nums {
		if i%k == 0 {
			maxLeft[i] = v
		} else {
			maxLeft[i] = Max(maxLeft[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if (i+1)%k == 0 || i == n-1 {
			maxRight[i] = nums[i]
		} else {
			maxRight[i] = Max(maxRight[i+1], nums[i])
		}
	}
	res := make([]int, n-k+1)
	for i := range res {
		res[i] = Max(maxRight[i], maxLeft[i+k-1])
	}
	return res
}
