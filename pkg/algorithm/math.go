package algorithm

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
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//BinarySearch 二分查找
func BinarySearch(n int, fn func(l, r, mid int) bool) int {
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if !fn(l, r, mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
