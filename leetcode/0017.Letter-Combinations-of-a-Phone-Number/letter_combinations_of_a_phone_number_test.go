package leetcode

import (
	"testing"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

func TestProblem17(t *testing.T) {

	qs := []question17{

		{
			para17{"23"},
			ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},

		{
			para17{""},
			ans17{[]string{}},
		},

		{
			para17{"2"},
			ans17{[]string{"a", "b", "c"}},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans17, q.para17
		t.Logf("[input]:%v		[output]:%v\n", p, letterCombinations(p.s))
		t.Logf("[input]:%v		[output]:%v\n", p, letterCombinationsBT(p.s))
	}

	t.Logf("------------------------------END------------------------------\n")

}
