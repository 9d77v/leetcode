package main

/*
题目：
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

限制：

0 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
*/

/*
方法一：归并排序
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：116 ms	内存消耗：8.2 MB
*/
func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	return mergeSort(nums, 0, n-1, make([]int, n))
}

func mergeSort(nums []int, left, right int, tmpArr []int) int {
	if left >= right {
		return 0
	}
	middle := int(uint(left+right) >> 1)
	cnt := mergeSort(nums, left, middle, tmpArr) + mergeSort(nums, middle+1, right, tmpArr)
	return merge(nums, left, right, middle, cnt, tmpArr)
}

func merge(nums []int, left, right, middle, cnt int, tmpArr []int) int {
	for i := left; i <= right; i++ {
		tmpArr[i] = nums[i]
	}
	i, j := left, middle+1
	for k := left; k <= right; k++ {
		if i > middle {
			nums[k] = tmpArr[j]
			j++
		} else if j > right {
			nums[k] = tmpArr[i]
			i++
		} else if tmpArr[i] > tmpArr[j] {
			nums[k] = tmpArr[j]
			j++
			cnt += middle - i + 1
		} else {
			nums[k] = tmpArr[i]
			i++
		}
	}
	return cnt
}
