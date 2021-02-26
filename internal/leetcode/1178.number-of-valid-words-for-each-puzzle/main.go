package main

import (
	"math/bits"
	"sort"
)

/*
题目：猜字谜
外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就可以算作谜底：

单词 word 中包含谜面 puzzle 的第一个字母。
单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及 "based"（其中的 "s" 没有出现在谜面中）。
返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words 中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。

提示：

1 <= words.length <= 10^5
4 <= words[i].length <= 50
1 <= puzzles.length <= 10^4
puzzles[i].length == 7
words[i][j], puzzles[i][j] 都是小写英文字母。
每个 puzzles[i] 所包含的字符都不重复。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle
*/

/*
方法一：对于每一个puzzle都遍历所有words
时间复杂度：О(n^3)
空间复杂度：О(n)
运行时间：超出时间限制
*/
func findNumOfValidWords(words []string, puzzles []string) []int {
	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		result[i] = getCount(puzzle, words)
	}
	return result
}

func getCount(puzzle string, words []string) int {
	arr := [26]int{}
	for i := range puzzle {
		arr[puzzle[i]-'a']++
	}
	count := 0
	for _, word := range words {
		flag := true
		hasFirst := false
		for i := 0; i < len(word); i++ {
			if arr[word[i]-'a'] == 0 {
				flag = false
				break
			}
			if word[i] == puzzle[0] {
				hasFirst = true
			}
		}
		if flag && hasFirst {
			count++
		}
	}
	return count
}

/*
方法二：二进制状态压缩
时间复杂度：О(m∣w∣+n2^∣p∣)
空间复杂度：О(m)
枚举子集方法一：运行时间：80 ms	内存消耗：13.6 MB
枚举子集方法二：运行时间：68 ms	内存消耗：13.6 MB
*/
func findNumOfValidWordsFunc2(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := map[uint]int{}
	for _, s := range words {
		var bitMask uint
		for _, ch := range s {
			bitMask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(bitMask) <= puzzleLength {
			cnt[bitMask]++
		}
	}
	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		var first uint = 1 << (puzzle[0] - 'a')
		//枚举子集方法一
		// for choose := 0; choose < 1<<(puzzleLength-1); choose++ {
		// 	var bitMask uint
		// 	for j := 0; j < puzzleLength-1; j++ {
		// 		if choose>>j&1 == 1 {
		// 			bitMask |= 1 << (puzzle[j+1] - 'a')
		// 		}
		// 	}
		// 	result[i] += cnt[bitMask|first]
		// }
		//枚举子集方法二
		var bitMask uint
		for _, ch := range puzzle[1:] {
			bitMask |= 1 << (ch - 'a')
		}
		subset := bitMask
		for {
			result[i] += cnt[subset|first]
			subset = (subset - 1) & bitMask
			if subset == bitMask {
				break
			}
		}
	}
	return result
}

/*
方法三：字典树
时间复杂度：О(m∣w∣log∣w∣+n2^∣p∣)
空间复杂度：О(m∣w∣+∣p∣)
运行时间：144 ms	内存消耗：38.8 MB
*/
func findNumOfValidWordsFunc3(words []string, puzzles []string) []int {
	root := new(trieNode)
	for _, word := range words {
		//将word中的字母按照字典序排序并去重
		w := []byte(word)
		sort.Slice(w, func(i, j int) bool { return w[i] < w[j] })
		i := 0
		for _, ch := range w[1:] {
			if w[i] != ch {
				i++
				w[i] = ch
			}
		}
		w = w[:i+1]

		//加入字典树中
		cur := root
		for _, ch := range w {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = new(trieNode)
			}
			cur = cur.son[ch]
		}
		cur.cnt++
	}
	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		pz := []byte(puzzle)
		first := pz[0]
		sort.Slice(pz, func(i, j int) bool { return pz[i] < pz[j] })

		var find func(int, *trieNode) int
		find = func(pos int, cur *trieNode) int {
			if cur == nil {
				return 0
			}
			if pos == len(pz) {
				return cur.cnt
			}

			res := find(pos+1, cur.son[pz[pos]-'a'])

			if pz[pos] != first {
				res += find(pos+1, cur)
			}
			return res
		}
		result[i] = find(0, root)
	}
	return result
}

type trieNode struct {
	son [26]*trieNode
	cnt int
}
