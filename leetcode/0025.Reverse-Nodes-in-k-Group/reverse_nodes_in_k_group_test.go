package leetcode

import (
	"testing"

	"github.com/motingjun/leetcode-go/structures"
)

type question25 struct {
	para25
	ans25
}

// para 是参数
// one 代表第一个参数
type para25 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans25 struct {
	one []int
}

func Test_Problem25(t *testing.T) {

	qs := []question25{

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			ans25{[]int{3, 2, 1, 4, 5}},
		},

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			ans25{[]int{1, 2, 3, 4, 5}},
		},
	}

	t.Logf("-----------------------LeetCode Problem------------------------\n")

	for _, q := range qs {
		_, p := q.ans25, q.para25
		t.Logf("[input]:%v		[output]:%v\n", p, structures.List2Ints(reverseKGroup(structures.Ints2List(p.one), p.two)))
	}

	t.Logf("------------------------------END------------------------------\n")
}
