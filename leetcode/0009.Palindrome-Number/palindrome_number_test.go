package leetcode

import (
	"testing"
)

type question9 struct {
	para9
	ans9
}

// para 是参数
// one 代表第一个参数
type para9 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans9 struct {
	one bool
}

func TestProblem9(t *testing.T) {

	qs := []question9{

		{
			para9{121},
			ans9{true},
		},

		{
			para9{-121},
			ans9{false},
		},

		{
			para9{10},
			ans9{false},
		},

		{
			para9{123},
			ans9{false},
		},

		{
			para9{0},
			ans9{true},
		},

		{
			para9{1534236469},
			ans9{false},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans9, q.para9
		t.Logf("[input]:%v		[output]:%v\n", p, isPalindrome(p.one))
		t.Logf("[input]:%v		[output]:%v\n", p, isPalindrome1(p.one))
	}

	t.Logf("------------------------------END------------------------------\n")
}
