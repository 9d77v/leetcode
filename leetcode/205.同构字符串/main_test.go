package main

import "testing"

type args struct {
	s string
	t string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"test ok", args{"egg", "add"}, true},
	{"test fail", args{"foo", "tar"}, false},
	{"test ok", args{"paper", "title"}, true},
	{"test ok", args{"ab", "aa"}, false},
	{"test ok", args{"bar", "foo"}, false},
}

func Test_isIsomorphic(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
