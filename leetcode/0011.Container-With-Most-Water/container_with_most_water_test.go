package leetcode

import (
	"testing"
)

type question11 struct {
	para11
	ans11
}

// para 是参数
// one 代表第一个参数
type para11 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	one int
}

func TestProblem11(t *testing.T) {

	qs := []question11{

		{
			para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans11{49},
		},

		{
			para11{[]int{1, 1}},
			ans11{1},
		},

		{
			para11{[]int{4, 3, 2, 1, 4}},
			ans11{16},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans11, q.para11
		t.Logf("[input]:%v		[output]:%v\n", p, maxArea(p.one))
	}

	t.Logf("------------------------------END------------------------------\n")
}
