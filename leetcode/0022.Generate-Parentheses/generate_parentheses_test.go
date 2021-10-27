package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "generateParenthesis",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findGenerateParenthesis(t *testing.T) {
	type args struct {
		lindex int
		rindex int
		str    string
		res    *[]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				lindex: 1,
				rindex: 1,
				str:    "(",
				res:    &[]string{"()"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findGenerateParenthesis(tt.args.lindex, tt.args.rindex, tt.args.str, tt.args.res)
		})
	}
}
