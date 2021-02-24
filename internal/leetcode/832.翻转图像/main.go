package main

/*
题目：
给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。

水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。

反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

提示：

1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flipping-an-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法一：模拟
时间复杂度：О(n^2)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：2.9 MB
*/
func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		reverse(A[i])
	}
	for i := range A {
		for j := range A[i] {
			A[i][j] ^= 1
		}
	}
	return A
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

/*
方法一：模拟+优化
时间复杂度：О(n^2)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：2.8 MB
*/
func flipAndInvertImageFunc2(A [][]int) [][]int {
	for i := range A {
		reverse2(A[i])
	}
	return A
}

func reverse2(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		if nums[i] == nums[n-1-i] {
			nums[i] ^= 1
			nums[n-1-i] ^= 1
		}
	}
	if n&1 == 1 {
		nums[(n-1)>>1] ^= 1
	}
}
