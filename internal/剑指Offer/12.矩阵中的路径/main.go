package main

/*
题目：
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

提示：
1 <= board.length <= 200
1 <= board[i].length <= 200

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
*/

/*
方法一：回溯法
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：4 ms	内存消耗：3.2 MB
*/
func exist(board [][]byte, word string) bool {
	n := len(board)
	if n == 0 {
		return false
	}
	m := len(board[0])
	if m == 0 || word == "" {
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = ' '
	if dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}
