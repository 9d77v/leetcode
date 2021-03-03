package algorithm

import (
	"math/rand"
	"time"
)

//Max return the maximum number
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Min return the minimum number
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//MaxArr return the maximum number
func MaxArr(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

//MinArr return the minimum number
func MinArr(a ...int) int {
	min := a[0]
	for _, v := range a[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

//Gcd 最大公约数
func Gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//Abs 绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Median 中位数
func Median(arr []int) float64 {
	n := len(arr)
	if n&1 == 1 {
		return float64(arr[n>>1])
	}
	return float64(arr[n>>1]+arr[n>>1-1]) / 2
}

//Reverse 反转数组
func Reverse(a []int) {
	for i, n := 0, len(a); i < n>>1; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//QuickSort 快排
func QuickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	qsort(nums, 0, n-1)
}

func qsort(nums []int, left, right int) {
	i := RandomPartition(nums, left, right)
	if i > left {
		qsort(nums, left, i-1)
	}
	if i < right {
		qsort(nums, i+1, right)
	}
}

//RandomPartition 随机分区
func RandomPartition(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	picked := left + rand.Int()%(right-left+1)
	nums[picked], nums[left] = nums[left], nums[picked]
	return Partition(nums, left, right)
}

//Partition 分区
func Partition(nums []int, left, right int) int {
	pivot, i := nums[left], left
	for i < right {
		for nums[right] >= pivot && i < right {
			right--
		}
		for nums[i] <= pivot && i < right {
			i++
		}
		nums[i], nums[right] = nums[right], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

//QuickSelect 快排变形求第k大
func QuickSelect(nums []int, k, left, right int) {
	i := RandomPartition(nums, left, right)
	if i > k {
		QuickSelect(nums, k, left, i-1)
	} else if i < k {
		QuickSelect(nums, k, i+1, right)
	}
}

//MergeSort 归并排序
func MergeSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	mergeSort(nums, 0, n-1, make([]int, n))
}

func mergeSort(nums []int, left, right int, tmpArr []int) {
	if left >= right {
		return
	}
	middle := int(uint(left+right) >> 1)
	mergeSort(nums, left, middle, tmpArr)
	mergeSort(nums, middle+1, right, tmpArr)
	merge(nums, left, right, middle, tmpArr)
}

func merge(nums []int, left, right, middle int, tmpArr []int) {
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
		} else {
			nums[k] = tmpArr[i]
			i++
		}
	}
}
