package main

/*
题目：在一个长度为 n+1 的数组 nums 里的所有数字都在 1～n 的范围内,所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的数组。
*/

/*
方法一：使用辅助数组
时间复杂度：О(n)
空间复杂度：О(n)
*/
func findRepeatNumberFunc1(nums []int) int {
	arr := make([]int, len(nums)+1)
	for _, num := range nums {
		if arr[num] == 1 {
			return num
		}
		arr[num]++
	}
	return -1
}

/*
方法二：二分查找
时间复杂度：О(n㏒n)
空间复杂度：О(1)
*/
func findRepeatNumberFunc2(nums []int) int {
	start := 1
	end := len(nums) - 1
	for end >= start {
		middle := (end-start)/2 + start
		count := countRange(nums, start, middle)
		if end == start {
			if count > 1 {
				return start
			}
			break
		}
		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}

	}
	return -1
}

func countRange(nums []int, start, end int) int {
	count := 0
	for _, num := range nums {
		if num >= start && num <= end {
			count++
		}
	}
	return count
}
