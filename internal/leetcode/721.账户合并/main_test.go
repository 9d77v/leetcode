package main

import (
	"reflect"
	"testing"
)

type args struct {
	accounts [][]string
}

var tests = []struct {
	name string
	args args
	want [][]string
}{
	{"should be equal", args{[][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}},
		[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}}},
}

func Test_accountsMerge(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
