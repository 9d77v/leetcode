package main

/*
题目：
假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。

注意:

数组内已种好的花不会违反种植规则。
输入的数组长度范围为 [1, 20000]。
n 是非负整数，且不会超过输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/can-place-flowers
*/

/*
方法一：遍历
时间复杂度：О(m)
空间复杂度：О(1)
运行时间：16 ms	内存消耗：6 MB
*/
func canPlaceFlowersFunc1(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	m := len(flowerbed)
	count := 0
	for i := 0; i < m; i++ {
		if flowerbed[i] == 0 && (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 == m || flowerbed[i+1] == 0) {
			count++
			if count == n {
				return true
			}
			flowerbed[i] = 1
		}
		if flowerbed[i] == 1 {
			i++
		}
	}
	return false
}
