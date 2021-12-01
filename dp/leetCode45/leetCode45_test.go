package leetcode45

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
			{1},
			{2,3,1,1,4},
			{2,3,0,1,4},
			{2,3,1,1,4,2,4,2,1,3,5,3,1,3,1},
		},
		want: []int{0,2,2,5},
	}

	for i:=0; i<len(j.in); i++ {
		j.got = append(j.got,jump(j.in[i]))
	}

	if !reflect.DeepEqual(j.got, j.want) {
		t.Errorf("expected:%v, got:%v", j.want, j.got)
	}
}
