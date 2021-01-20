package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
每年，政府都会公布一万个最常见的婴儿名字和它们出现的频率，也就是同名婴儿的数量。有些名字有多种拼法，例如，John 和 Jon 本质上是相同的名字，但被当成了两个名字公布出来。给定两个列表，一个是名字及对应的频率，另一个是本质相同的名字对。设计一个算法打印出每个真实名字的实际频率。注意，如果 John 和 Jon 是相同的，并且 Jon 和 Johnny 相同，则 John 与 Johnny 也相同，即它们有传递和对称性。

在结果列表中，选择 字典序最小 的名字作为真实名字。

提示：
names.length <= 100000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/baby-names-lcci
*/

/*
方法一：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：200 ms	内存消耗：8.3 MB
*/
func trulyMostPopular(names []string, synonyms []string) []string {
	if len(names) <= 1 {
		return names
	}
	idMap := make(map[string]int, 0)
	var uf UnionFind = NewMapUnionFindWithRank(RankSize)
	for _, synonym := range synonyms {
		a, b := getSynonym(synonym)
		if _, ok := idMap[a]; !ok {
			idMap[a] = len(idMap)
		}
		if _, ok := idMap[b]; !ok {
			idMap[b] = len(idMap)
		}
		uf.Union(idMap[a], idMap[b])
	}

	type tmp struct {
		index int
		name  string
	}
	rootMap := make(map[int]*tmp, 0)
	result := []string{}
	for _, v := range names {
		name, count := getNameAndCount(v)
		if _, ok := idMap[name]; !ok {
			result = append(result, v)
		} else {
			rootID := uf.Find(idMap[name])
			if _, ok := rootMap[rootID]; !ok {
				rootMap[rootID] = &tmp{
					index: len(result),
					name:  name,
				}
				result = append(result, v)
			} else {
				if name < rootMap[rootID].name {
					rootMap[rootID].name = name
				}
				rootIndex := rootMap[rootID].index
				vName, vCount := getNameAndCount(result[rootIndex])
				vName, vCount = rootMap[rootID].name, vCount+count
				result[rootIndex] = formatName(vName, vCount)
			}
		}
	}
	return result
}

func getSynonym(s string) (string, string) {
	arr := strings.Split(s[1:len(s)-1], ",")
	return arr[0], arr[1]
}

func getNameAndCount(s string) (string, int) {
	arr := strings.Split(s[:len(s)-1], "(")
	count, _ := strconv.Atoi(arr[1])
	return arr[0], count
}

func formatName(name string, count int) string {
	return fmt.Sprintf("%s(%d)", name, count)
}
