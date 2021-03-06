package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：买卖股票的最佳时机含手续费
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

注意:
0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法一：动态规划
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：108 ms	内存消耗：7.1 MB
*/
func maxProfitFunc1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0] - fee}
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：108 ms	内存消耗：7.8 MB
*/
func maxProfitFunc2(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	profit0, profit1 := 0, -prices[0]-fee
	for i := 1; i < n; i++ {
		profit0 = Max(profit0, profit1+prices[i])
		profit1 = Max(profit1, profit0-prices[i]-fee)
	}
	return profit0
}
