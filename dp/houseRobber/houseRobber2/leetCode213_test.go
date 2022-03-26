package houserobber2

import "testing"

func TestRob(t *testing.T) {
	in := []int{1,2,3,1,2,3,4,2,1,3,2}
	got := rob(in)
	want := 13
	if got != want {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}