package main

import (
	"reflect"
	"testing"
)

type args struct {
	words   []string
	puzzles []string
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
		[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}},
		[]int{1, 1, 3, 2, 4, 0}},
	{"should be equal", args{[]string{"apple", "pleas", "please"},
		[]string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}},
		[]int{0, 1, 3, 2, 0}},
}

func Test_findNumOfValidWords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumOfValidWordsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWordsFunc2(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWordsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumOfValidWordsFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWordsFunc3(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWordsFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
