package leetcode62

import (
	"reflect"
	"testing"
)

type jumpTest struct {
	in   [][]int
	got  []int
	want []int
}

func TestJump(t *testing.T) {
	j := jumpTest{
		in: [][]int{
			{1,1},
			{3,7},
			{3,2},
			{7,3},
			{3,3},
			{100,100},
		},
		want: []int{1,28,3,28,6,4631081169483718960},
	}

	for i:=0; i<len(j.in); i++ {
		j.got = append(j.got,uniquePaths(j.in[i][0], j.in[i][1]))
	}

	if !reflect.DeepEqual(j.got, j.want) {
		t.Errorf("expected:%v, got:%v", j.want, j.got)
	}
}