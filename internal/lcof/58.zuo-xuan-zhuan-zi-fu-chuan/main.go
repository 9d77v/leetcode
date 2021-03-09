package main

/*
题目：左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof
*/

/*
方法一：三次旋转
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.3 MB
*/
func reverseLeftWords(s string, n int) string {
	arr, m := []byte(s), len(s)
	reverseString(arr, 0, m-1)
	reverseString(arr, 0, m-n-1)
	reverseString(arr, m-n, m-1)
	return string(arr)
}

func reverseString(arr []byte, left, right int) {
	for i := left; i < (left+right+1)>>1; i++ {
		arr[i], arr[right+left-i] = arr[right+left-i], arr[i]
	}
	return
}

/*
方法二：拼接字符串
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.1 MB
*/
func reverseLeftWordsFunc2(s string, n int) string {
	return s[n:] + s[:n]
}
