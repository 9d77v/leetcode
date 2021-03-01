package main

/*
题目：
给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

提示：

0 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= i <= j < nums.length
最多调用 104 次 sumRange 方法

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-immutable
*/

/*
方法一: 前缀和
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：32 ms	内存消耗：9.4 MB
*/

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{
		preSum: make([]int, len(nums)),
	}
	for i, num := range nums {
		if i > 0 {
			arr.preSum[i] = arr.preSum[i-1] + num
		} else {
			arr.preSum[i] = num
		}
	}
	return arr
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.preSum[j] - this.preSum[i-1]
	}
	return this.preSum[j]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
