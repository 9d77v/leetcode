package main

import (
	. "github.com/9d77v/leetcode/lib"
)

/*
题目：
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
*/

/*
方法一：动态规划
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.4 MB
*/
func maxProfitFunc1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		profit0 := 0
		if i >= 2 {
			profit0 = dp[i-2][0]
		}
		dp[i][1] = Max(dp[i-1][1], profit0-prices[i])
	}
	return dp[n-1][0]
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.3 MB
*/
func maxProfitFunc2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	prevProfit0, profit0, profit1 := 0, 0, -prices[0]
	for i := 1; i < n; i++ {
		nextProfit0 := Max(profit0, profit1+prices[i])
		nextProfit1 := Max(profit1, prevProfit0-prices[i])
		prevProfit0 = profit0
		profit0 = nextProfit0
		profit1 = nextProfit1
	}
	return profit0
}
