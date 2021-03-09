package main

/*
题目：
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
*/

/*
方法一：递归
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func removeDuplicates(S string) string {
	newSArr := []byte{}
	hasCut := false
	if len(S) == 1 {
		return S
	}
	for i := 0; i < len(S)-1; i++ {
		if S[i] == S[i+1] {
			hasCut = true
			i++
			if i == len(S)-2 {
				newSArr = append(newSArr, S[i+1])
			}
		} else {
			newSArr = append(newSArr, S[i])
			if i == len(S)-2 {
				newSArr = append(newSArr, S[i+1])
			}
		}
	}
	if !hasCut {
		return string(newSArr)
	}
	return removeDuplicates(string(newSArr))
}

/*
方法二：栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：6.3 MB
*/
func removeDuplicatesFunc2(S string) string {
	stack := []byte{}
	for i := range S {
		if len(stack) != 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
