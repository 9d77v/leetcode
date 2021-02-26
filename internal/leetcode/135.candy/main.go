package main

/*
题目：分发糖果
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/candy
*/

/*
方法一：两次遍历
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：20 ms	内存消耗：6.2 MB
*/
func candyFunc1(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	sum := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && arr[i] <= arr[i+1] {
			arr[i] = arr[i+1] + 1
		}
		sum += arr[i]
	}
	return sum
}

/*
方法二：常数空间遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：16 ms	内存消耗：6.1 MB
*/
func candyFunc2(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	sum, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			sum += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			sum += dec
			pre = 1
		}
	}
	return sum
}
