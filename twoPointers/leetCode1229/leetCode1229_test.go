package leetcode1229

import (
	"reflect"
	"testing"
)

func TestMinAvailableDuration(t *testing.T) {
	slots1 := [][]int{{10,50},{60,120},{140,210}}
	slots2 := [][]int{{0,15},{60,70}}
	duration := 8
	want := []int{60,68}
	got := minAvailableDuration(slots1,slots2,duration)

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expect:%v, got:%v",want, got)
	}
}