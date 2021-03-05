package main

/*
题目：和为s的连续正数序列

输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

来源：力扣（LeetCode）
https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
*/

/*
方法一：枚举+暴力
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：112 ms	内存消耗：6.9 MB
*/
func findContinuousSequence(target int) [][]int {
	result, i := [][]int{}, 1
	for {
		arr, sum := []int{}, 0
		for k := i; sum < target; k++ {
			sum += k
			arr = append(arr, k)
		}
		if sum == target {
			if len(arr) < 2 {
				return result
			}
			result = append(result, arr)
		}
		if len(arr) < 2 {
			return result
		}
		i++
	}
}

/*
方法二：枚举+双指针
时间复杂度：О(target)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.2 MB
*/
func findContinuousSequenceFunc2(target int) [][]int {
	result := [][]int{}
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			arr := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				arr[i-l] = i
			}
			result = append(result, arr)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return result
}
