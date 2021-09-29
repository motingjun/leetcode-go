package leetcode

import (
	"testing"
)

type question6 struct {
	para6
	ans6
}

// para 是参数
// s 代表字符串, numRows 代表行数
type para6 struct {
	s       string
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans6 struct {
	one string
}

func TestProblem6(t *testing.T) {

	qs := []question6{

		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},

		{
			para6{"PAYPALISHIRING", 4},
			ans6{"PINALSIGYAHRPI"},
		},

		{
			para6{"A", 1},
			ans6{"A"},
		},
	}

	t.Logf("-----------------------LeetCode Problem 4-----------------------\n")

	for _, q := range qs {
		_, p := q.ans6, q.para6
		t.Logf("[input]:%v		[output]:%v\n", p, convert(p.s, p.numRows))
	}

	t.Logf("------------------------------END------------------------------\n")
}
