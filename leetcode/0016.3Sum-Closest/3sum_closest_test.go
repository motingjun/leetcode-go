package leetcode

import (
	"testing"
)

type question16 struct {
	para16
	ans16
}

// para 是参数
// a 代表第一个参数
// target 代表目标值
type para16 struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans16 struct {
	one int
}

func TestProblem16(t *testing.T) {

	qs := []question16{

		{
			para16{[]int{-1, 0, 1, 1, 55}, 3},
			ans16{2},
		},

		{
			para16{[]int{0, 0, 0}, 1},
			ans16{0},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},

		{
			para16{[]int{1, 1, -1}, 0},
			ans16{1},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},

		{
			para16{[]int{-1, 2}, 1},
			ans16{0},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans16, q.para16
		t.Logf("[input]:%v		[output]:%v\n", p, threeSumClosest(p.a, p.target))
		t.Logf("[input]:%v		[output]:%v\n", p, threeSumClosest1(p.a, p.target))
	}

	t.Logf("------------------------------END------------------------------\n")
}
