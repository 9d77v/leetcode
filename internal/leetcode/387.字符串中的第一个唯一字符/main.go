package main

/*
题目：给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
提示：你可以假定该字符串只包含小写字母。。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/

/*
方法一：使用哈希表存储频数
时间复杂度： О(n),2次О(n)遍历
空间复杂度：О(|Σ|)，|Σ|≤26
运行时间：8 ms	内存消耗：5.2 MB
*/
func firstUniqCharFunc1(s string) int {
	arr := [26]int{}
	for _, ch := range s {
		arr[ch-'a']++
	}
	for i, ch := range s {
		if arr[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

/*
方法二：使用哈希表存储索引
时间复杂度： О(n)，第一次О(n)，第二次О(|Σ|)
空间复杂度：О(|Σ|)，|Σ|≤26
运行时间：8 ms	内存消耗：5.2 MB
*/
func firstUniqCharFunc2(s string) int {
	arr := [26]int{}
	n := len(s)
	for i := range arr {
		arr[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if arr[ch] == n {
			arr[ch] = i
		} else {
			arr[ch] = n + 1
		}
	}
	ans := n
	for _, p := range arr {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}

/*
方法三：队列
时间复杂度： О(n)，遍历字符串О(n)，维护队列时间复杂度О(|Σ|)
空间复杂度：О(|Σ|)，|Σ|≤26 ，一个索引数组，一个队列
运行时间：8 ms	内存消耗：5.4 MB
*/
func firstUniqCharFunc3(s string) int {
	type pair struct {
		ch  byte
		pos int
	}

	n := len(s)
	arr := [26]int{}
	for i := range arr {
		arr[i] = n
	}
	queue := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if arr[ch] == n {
			arr[ch] = i
			queue = append(queue, pair{ch, i})
		} else {
			arr[ch] = n + 1
			for len(queue) > 0 && arr[int(queue[0].ch)] == n+1 {
				queue = queue[1:]
			}
		}
	}
	if len(queue) > 0 {
		return queue[0].pos
	}
	return -1
}
