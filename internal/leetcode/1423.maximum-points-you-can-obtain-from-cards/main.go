package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：可获得的最大点数
几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。

每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。

你的点数就是你拿到手中的所有卡牌的点数之和。

给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。

提示：

1 <= cardPoints.length <= 10^5
1 <= cardPoints[i] <= 10^4
1 <= k <= cardPoints.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards
*/

/*
方法一：双指针
时间复杂度：О(2k)
空间复杂度：О(1)
运行时间：64 ms	内存消耗：8.5 MB
*/
func maxScore(cardPoints []int, k int) int {
	windowSum, sum, n := 0, 0, len(cardPoints)
	for left := 0; left < k; left++ {
		windowSum += cardPoints[left]
	}
	sum = Max(sum, windowSum)
	for left := k - 1; left > -1; left-- {
		right := n - k + left
		windowSum += cardPoints[right] - cardPoints[left]
		sum = Max(windowSum, sum)
	}
	return sum
}

/*
方法二：双指针,滑动窗口求最小值
时间复杂度：О(n-k)
空间复杂度：О(1)
运行时间：64 ms	内存消耗：8.5 MB
*/
func maxScoreFunc2(cardPoints []int, k int) int {
	n := len(cardPoints)
	left, windowSize, windowSum, sum, totalSum := 0, n-k, 0, 0, 0
	for right, point := range cardPoints {
		totalSum += point
		windowSum += point
		if right == windowSize-1 {
			sum = windowSum
		} else if right > windowSize-1 {
			windowSum -= cardPoints[left]
			left++
			sum = Min(sum, windowSum)
		}
	}
	return totalSum - sum
}
