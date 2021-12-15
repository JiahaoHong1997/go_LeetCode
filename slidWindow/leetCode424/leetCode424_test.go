package leetcode424

import (

	"testing"
)


func TestCharacterReplacementHeap(t *testing.T) {
	s := "AABABBA"
	k := 1
	want := 4

	got := characterReplacementHeap(s,k)
	if want != got {
		t.Errorf("expected:%v, got:%v",want,got)
	}
}