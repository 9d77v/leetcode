package main

/*
题目：最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：O(k)
运行时间：284 ms	内存消耗：13 MB
*/
func minWindowFunc1(s string, t string) string {
	left, right, chMap, tMap := 0, 0, make(map[byte]int, 52), make(map[byte]int)
	for _, ch := range t {
		tMap[byte(ch)]++
	}
	for ; right < len(s); right++ {
		if _, ok := tMap[s[right]]; ok {
			chMap[s[right]]++
		}
		if len(chMap) == 0 {
			left++
		}
		if contains(chMap, tMap) {
			break
		}
	}
	if !contains(chMap, tMap) {
		return ""
	}
	for i := left; i <= right; i++ {
		if _, ok := chMap[s[left]]; !ok {
			left++
		} else if chMap[s[left]] > tMap[s[left]] {
			chMap[s[left]]--
			left++
		} else {
			break
		}
	}
	right++
	result := s[left:right]
	for ; right < len(s); right++ {
		if _, ok := chMap[s[left]]; ok {
			chMap[s[left]]--
			if chMap[s[left]] == 0 {
				delete(chMap, s[left])
			}
		}
		if _, ok := tMap[s[right]]; ok {
			chMap[s[right]]++
		}
		left++
		if contains(chMap, tMap) {
			for i := left; i <= right; i++ {
				if _, ok := chMap[s[left]]; !ok {
					left++
				} else if chMap[s[left]] > tMap[s[left]] {
					chMap[s[left]]--
					left++
				} else {
					break
				}
			}
			result = s[left : right+1]
		}
	}
	return result
}

func contains(chMap, tMap map[byte]int) bool {
	ok := true
	for k, v := range tMap {
		if chMap[k] < v {
			ok = false
		}
	}
	return ok
}
