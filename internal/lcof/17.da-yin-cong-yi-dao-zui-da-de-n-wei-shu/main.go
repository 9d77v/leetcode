package main

import "strconv"

/*
题目：打印从一到最大的n位数
*/

/*
方法一：循环
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：7 MB
*/
func printNumbersFunc1(n int) []int {
	k := 1
	for n > 0 {
		k *= 10
		n--
	}
	res := make([]int, k-1)
	for i := 1; i < k; i++ {
		res[i-1] = i
	}
	return res
}

/*
方法一：dfs
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：18 ms	内存消耗：8 MB
*/

func printNumbersFunc2(n int) []int {
	res := make([]int, 0)
	dfs(&res, 0, n, make([]byte, n))
	return res
}

var nums = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func dfs(res *[]int, depth, n int, digits []byte) {
	if depth == n {
		i := 0
		for ; i < len(digits); i++ {
			if digits[i] != '0' {
				break
			}
		}
		if i == n {
			return
		}
		x, _ := strconv.Atoi(string(digits[i:]))
		*res = append(*res, x)
		return
	}
	for _, v := range nums {
		digits[depth] = v
		dfs(res, depth+1, n, digits)
	}
}
