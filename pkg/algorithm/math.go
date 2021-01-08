package algorithm

//Max return the maximum number
func Max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

//Min return the minimum number
func Min(a ...int) int {
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

//Reverse 反转数组
func Reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
