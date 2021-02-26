package main

/*
题目：同构字符串
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

说明:
你可以假设 s 和 t 具有相同的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/isomorphic-strings
*/

/*
方法一：哈希表
时间复杂度：О(n)
空间复杂度：O(|Σ|)
运行时间：4 ms	内存消耗：2.5 MB
*/
func isIsomorphic(s string, t string) bool {
	mapA := make(map[byte]byte)
	mapB := make(map[byte]bool)
	for i := range s {
		x, y := s[i], t[i]
		if mapA[x] == 0 {
			if !mapB[y] {
				mapA[x] = y
				mapB[y] = true
				continue
			}
			return false
		}
		if mapA[x] != y {
			return false
		}
	}
	return true
}
