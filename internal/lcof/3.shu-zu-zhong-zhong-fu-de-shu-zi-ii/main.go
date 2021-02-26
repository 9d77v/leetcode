package main

/*
题目：数组中重复的数字II
在一个长度为 n+1 的数组 nums 里的所有数字都在 1～n 的范围内,所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的数组。
*/

/*
方法一：使用辅助hash
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
运行时间：8 ms	内存消耗：3.8 MB
*/
func findRepeatNumberFunc2(nums []int) (result int) {
	l := 1
	r := len(nums) - 1
	for r >= l {
		middle := (l + r) >> 1
		cnt := 0
		for _, num := range nums {
			if num <= middle {
				cnt++
			}
		}
		if cnt <= middle {
			l = middle + 1
		} else {
			r = middle - 1
			result = middle
		}
	}
	return
}

/*
方法三：快慢指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.8 MB
*/
func findRepeatNumberFunc3(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
