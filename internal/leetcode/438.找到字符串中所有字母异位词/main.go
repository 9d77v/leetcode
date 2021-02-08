package main

/*
题目：
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：20 ms	内存消耗：5 MB
*/
func findAnagrams(s string, p string) []int {
	chArr, m, result := [26]int{}, len(p), []int{}
	count := m
	for i := range p {
		chArr[p[i]-'a']++
	}
	for right := range s {
		if chArr[s[right]-'a'] > 0 {
			chArr[s[right]-'a']--
			if chArr[s[right]-'a'] >= 0 {
				count--
			}
		}
		left := right - m
		if left >= 0 {
			if chArr[s[left]-'a'] > 0 {
				chArr[s[left]-'a']++
				if chArr[s[left]-'a'] > 0 {
					count++
				}
			}
		}
		if count == 0 {
			result = append(result, left+1)
		}
	}
	return result
}
