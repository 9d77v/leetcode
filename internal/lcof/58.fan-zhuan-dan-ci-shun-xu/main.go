package main

import (
	"strings"
)

/*
题目：翻转单词顺序

输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

来源：力扣（LeetCode）
https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
*/

/*
方法一：字符串分割+倒序
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.5 MB
*/
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	newArr := []string{}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		}
		newArr = append(newArr, arr[i])
	}
	return strings.Join(newArr, " ")
}

/*
方法一：字符串提取单词+倒序
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.3 MB
*/
func reverseWordsFunc2(s string) string {
	arr := strings.Fields(s)
	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}

/*
方法三：双指针
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3.9 MB
*/
func reverseWordsFunc3(s string) string {
	arr := trimSpace([]byte(s))
	result := []byte{}
	i, j := len(arr)-1, len(arr)-1
	for i >= 0 {
		for i >= 0 && arr[i] != ' ' {
			i--
		}
		result = append(result, append([]byte(arr[i+1:j+1]), ' ')...)
		for i >= 0 && arr[i] == ' ' {
			i--
		}
		j = i
	}
	if len(result) == 0 {
		return ""
	}
	return string(result[:len(result)-1])
}

func trimSpace(arr []byte) []byte {
	start, end := 0, len(arr)-1
	for start <= end {
		if arr[start] != ' ' {
			break
		}
		start++
	}
	for end >= start {
		if arr[end] != ' ' {
			break
		}
		end--
	}
	return arr[start : end+1]
}
