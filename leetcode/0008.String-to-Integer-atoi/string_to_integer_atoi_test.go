package leetcode

import (
	"testing"
)

type question8 struct {
	para8
	ans8
}

// para 是参数
// one 代表第一个参数
type para8 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans8 struct {
	one int
}

func TestProblem8(t *testing.T) {

	qs := []question8{

		{
			para8{"42"},
			ans8{42},
		},

		{
			para8{"   -42"},
			ans8{-42},
		},

		{
			para8{"4193 with words"},
			ans8{4193},
		},

		{
			para8{"words and 987"},
			ans8{0},
		},

		{
			para8{"-91283472332"},
			ans8{-2147483648},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans8, q.para8
		t.Logf("[input]:%v		[output]:%v\n", p, myAtoi(p.s))
	}

	t.Logf("------------------------------END------------------------------\n")
}
