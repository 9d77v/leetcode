package main

import (
	"math"

	. "github.com/9d77v/leetcode/lib"
)

/*
题目：
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
0 <= k <= 10^9
0 <= prices.length <= 1000
0 <= prices[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
*/

/*
方法一：动态规划
时间复杂度：О(nmin(n,k))
空间复杂度：О(nmin(n,k))
运行时间：4 ms	内存消耗：4.6 MB
*/
func maxProfitFunc1(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= n/2 {
		return maxProfit1(prices)
	}
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 1; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/*
方法三：动态规划
时间复杂度：О(nmin(n,k))
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.2 MB
*/
func maxProfitFunc2(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= n/2 {
		return maxProfit2(prices)
	}
	dp := make([][2]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = Max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = Max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}

func maxProfit2(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}
	profit0, profit1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		profit0 = Max(profit0, profit1+prices[i])
		profit1 = Max(profit1, profit0-prices[i])
	}
	return profit0
}

/*
方法三：动态规划
时间复杂度：О(nmin(n,k))
空间复杂度：О(min(n,k))
运行时间：0 ms	内存消耗：2.2 MB
*/
func maxProfitFunc3(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = Min(k, n/2)
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}
	for i := 1; i < n; i++ {
		buy[0] = Max(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = Max(buy[j], sell[j]-prices[i])
			sell[j] = Max(sell[j], buy[j-1]+prices[i])
		}
	}
	return Max(sell...)
}
