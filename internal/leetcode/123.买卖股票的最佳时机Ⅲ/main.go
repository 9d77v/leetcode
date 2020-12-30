package main

import (
	"math"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
*/

/*
方法一：动态规划
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：188 ms	内存消耗：13.7 MB
*/
func maxProfitFunc1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3][2]int, n)
	dp[0] = [3][2]int{
		{0, math.MinInt64},
		{0, -prices[0]},
		{0, -prices[0]},
	}
	for i := 1; i < n; i++ {
		dp[i][2][0] = Max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = Max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
		dp[i][1][0] = Max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = Max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
	}
	return dp[n-1][2][0]
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：168 ms	内存消耗：8.8 MB
*/
func maxProfitFunc2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	profitOne0, profitOne1, profitTwo0, profitTwo1 := 0, -prices[0], 0, -prices[0]
	for i := 1; i < n; i++ {
		profitTwo0 = Max(profitTwo0, profitTwo1+prices[i])
		profitTwo1 = Max(profitTwo1, profitOne0-prices[i])
		profitOne0 = Max(profitOne0, profitOne1+prices[i])
		profitOne1 = Max(profitOne1, -prices[i])
	}
	return profitTwo0
}
