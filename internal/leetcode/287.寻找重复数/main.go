package main

/*
题目：
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

提示：

2 <= n <= 3 * 104
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
*/

/*
方法一：二分查找
时间复杂度：О(n㏒n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.8 MB
*/
func findRepeatNumberFunc1(nums []int) (result int) {
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
方法二：快慢指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.8 MB
*/
func findRepeatNumberFunc2(nums []int) int {
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
