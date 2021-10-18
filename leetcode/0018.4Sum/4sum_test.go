package leetcode

import "testing"

type question18 struct {
	para18
	ans18
}

// para 是参数
// a 代表第一个参数
// target 代表目标值
type para18 struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans18 struct {
	one [][]int
}

func TestProblem18(t *testing.T) {

	qs := []question18{

		{
			para18{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans18{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},

		{
			para18{[]int{2, 2, 2, 2, 2}, 8},
			ans18{[][]int{{2, 2, 2, 2}}},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans18, q.para18
		t.Logf("[input]:%v		[output]:%v\n", p, fourSum(p.a, p.target))
		t.Logf("[input]:%v		[output]:%v\n", p, fourSum1(p.a, p.target))
	}

	t.Logf("------------------------------END------------------------------\n")
}
