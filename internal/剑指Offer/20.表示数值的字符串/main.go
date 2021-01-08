package main

/*
题目：
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof
*/

/*
方法一：遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：2.3 MB
*/
func isNumberFunc1(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	start, end := 0, n-1
	for ; start < n; start++ {
		if s[start] != ' ' {
			break
		}
	}
	for ; end >= 0; end-- {
		if s[end] != ' ' {
			break
		}
	}
	ok, start := scanInteger(s, start, end)
	if start <= end && s[start] == '.' {
		start++
		ok2, j := scanUnsignInteger(s, start, end)
		ok, start = ok || ok2, j
	}
	if start <= end && (s[start] == 'e' || s[start] == 'E') {
		start++
		ok3, k := scanInteger(s, start, end)
		ok, start = ok3 && ok, k
	}
	return ok && start > end
}

func scanUnsignInteger(s string, start, end int) (bool, int) {
	cur := start
	for start <= end && s[start] >= '0' && s[start] <= '9' {
		start++
	}
	return start > cur, start
}

func scanInteger(s string, start, end int) (bool, int) {
	if start <= end && (s[start] == '+' || s[start] == '-') {
		start++
	}
	return scanUnsignInteger(s, start, end)
}

/*
方法二： 确定有限状态机
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：2.3 MB
*/
func isNumberFunc2(s string) bool {
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transferMap[state][typ]; !ok {
			return false
		}
		state = transferMap[state][typ]
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

var transferMap = map[State]map[CharType]State{
	STATE_INITIAL: {
		CHAR_SPACE:  STATE_INITIAL,
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		CHAR_SIGN:   STATE_INT_SIGN,
	},
	STATE_INT_SIGN: {
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_POINT:  STATE_POINT_WITHOUT_INT,
	},
	STATE_INTEGER: {
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_EXP:    STATE_EXP,
		CHAR_POINT:  STATE_POINT,
		CHAR_SPACE:  STATE_END,
	},
	STATE_POINT: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_EXP:    STATE_EXP,
		CHAR_SPACE:  STATE_END,
	},
	STATE_POINT_WITHOUT_INT: {
		CHAR_NUMBER: STATE_FRACTION,
	},
	STATE_FRACTION: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_EXP:    STATE_EXP,
		CHAR_SPACE:  STATE_END,
	},
	STATE_EXP: {
		CHAR_NUMBER: STATE_EXP_NUMBER,
		CHAR_SIGN:   STATE_EXP_SIGN,
	},
	STATE_EXP_SIGN: {
		CHAR_NUMBER: STATE_EXP_NUMBER,
	},
	STATE_EXP_NUMBER: {
		CHAR_NUMBER: STATE_EXP_NUMBER,
		CHAR_SPACE:  STATE_END,
	},
	STATE_END: {
		CHAR_SPACE: STATE_END,
	},
}
