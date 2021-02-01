package main

import "sort"

/*
题目：
爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。

因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）

返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。

如果有多个答案，你可以返回其中任何一个。保证答案存在。

提示：

1 <= A.length <= 10000
1 <= B.length <= 10000
1 <= A[i] <= 100000
1 <= B[i] <= 100000
保证爱丽丝与鲍勃的糖果总量不同。
答案肯定存在。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fair-candy-swap
*/

/*
方法一：遍历
时间复杂度：О(nm)
空间复杂度：О(1)
运行时间：288 ms	内存消耗：6.8 MB
*/
func fairCandySwapFunc1(A []int, B []int) []int {
	sumA, sumB := sum(A), sum(B)
	delta := (sumA - sumB) >> 1
	for _, a := range A {
		for _, b := range B {
			if a-b == delta {
				return []int{a, b}
			}
		}
	}
	return []int{}
}

func sum(A []int) int {
	count := 0
	for _, a := range A {
		count += a
	}
	return count
}

/*
方法二：哈希表
时间复杂度：О(n+m)
空间复杂度：О(n)
运行时间：64 ms	内存消耗：7.1 MB
*/
func fairCandySwapFunc2(A []int, B []int) []int {
	sumA := 0
	aMap := make(map[int]struct{}, len(A))
	for _, a := range A {
		sumA += a
		aMap[a] = struct{}{}
	}
	sumB := 0
	for _, b := range B {
		sumB += b
	}
	delta := (sumA - sumB) >> 1
	for _, b := range B {
		a := b + delta
		if _, ok := aMap[a]; ok {
			return []int{a, b}
		}
	}
	return []int{}
}

/*
方法三：排序+双指针
时间复杂度：О(nlogm)
空间复杂度：О(1)
运行时间：96 ms	内存消耗：6.8 MB
*/
func fairCandySwapFunc3(A []int, B []int) []int {
	sumA, sumB := sum(A), sum(B)
	delta := (sumA - sumB) >> 1
	sort.Ints(A)
	sort.Ints(B)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		if A[i]-B[j] == delta {
			return []int{A[i], B[j]}
		}
		if A[i]-B[j] > delta {
			j++
		} else {
			i++
		}
	}
	return []int{}
}
