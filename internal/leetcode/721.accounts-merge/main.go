package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：账户合并
给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。

现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。账户本身可以以任意顺序返回。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/accounts-merge
*/

/*
方法一：并查集(数组)
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：64 ms	内存消耗：8.7 MB
*/
func accountsMerge(accounts [][]string) [][]string {
	idStrMap, idMap := getIDMap(accounts)
	uf := NewArrayUnionFind(len(idMap))
	for _, account := range accounts {
		if len(account) < 2 {
			continue
		}
		firstIndex := idStrMap[account[1]]
		for i := 1; i < len(account); i++ {
			uf.Union(idStrMap[account[i]], firstIndex)
		}
	}
	dMap := make(map[int][]string, uf.Count())
	for i := 0; i < len(idMap); i++ {
		key := uf.Find(i)
		dMap[key] = append(dMap[key], idMap[i])
	}
	addedMap := make(map[int]bool, 0)
	result := make([][]string, 0)
	for _, account := range accounts {
		if len(account) < 2 {
			continue
		}
		rootID := uf.Find(idStrMap[account[1]])
		if _, ok := addedMap[rootID]; !ok {
			addedMap[rootID] = true
			newAccount := []string{account[0]}
			sort.Strings(dMap[rootID])
			newAccount = append(newAccount, dMap[rootID]...)
			result = append(result, newAccount)
		}
	}
	return result
}

func getIDMap(accounts [][]string) (map[string]int, map[int]string) {
	result1 := make(map[string]int, 0)
	result2 := make(map[int]string, 0)
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			if _, ok := result1[account[i]]; !ok {
				result1[account[i]] = len(result1)
				result2[len(result2)] = account[i]
			}
		}
	}
	return result1, result2
}
