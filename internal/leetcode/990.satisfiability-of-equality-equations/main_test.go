package main

import (
	"testing"
)

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be false", args{[]string{"a==b", "b!=a"}}, false},
		{"should be true", args{[]string{"b==a", "a==b"}}, true},
		{"should be true", args{[]string{"a==b", "b==c", "a==c"}}, true},
		{"should be false", args{[]string{"a==b", "b!=c", "c==a"}}, false},
		{"should be true", args{[]string{"c==c", "b==d", "x!=z"}}, true},
		{"should be false", args{[]string{"a!=a"}}, false},
		{"should be ture", args{[]string{"c==c", "f!=a", "f==b", "b==c"}}, true},
		{"should be true", args{[]string{"e==d", "e==a", "f!=d", "b!=c", "a==b"}}, true},
		{"should be false", args{[]string{"e!=c", "b!=b", "b!=a", "e==d"}}, false},
		{"should be false", args{[]string{"a!=f", "c!=a", "a!=d", "a==d", "f==c", "a!=c"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
