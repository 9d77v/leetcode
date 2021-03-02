package main

/*
题目：
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
*/

/*
方法一：哈希表
时间复杂度：О(n)
空间复杂度：О(C)
运行时间：4 ms	内存消耗：5.3 MB
*/
func firstUniqChar(s string) byte {
	chArr := [26]int{}
	for i := range s {
		chArr[s[i]-'a']++
	}
	for i := range s {
		if chArr[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
