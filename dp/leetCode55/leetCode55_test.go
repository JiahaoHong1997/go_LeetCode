package leetcode55

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {

	type canJumpTest struct {
		in   [][]int
		got  []bool
		want []bool
	}

	c := canJumpTest{
		in: [][]int{{2, 3, 1, 1, 4},
			{3, 2, 1, 0, 4},
			{0},
			{0, 3, 2, 4}},
	}
	for i := 0; i < len(c.in); i++ {
		c.got = append(c.got, canJump(c.in[i]))
	}
	c.want = []bool{true,
		false,
		true,
		false}

	if !reflect.DeepEqual(c.got, c.want) {
		t.Errorf("expected:%v, got:%v", c.want, c.got)
	}
}
