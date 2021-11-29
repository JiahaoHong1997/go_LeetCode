package leetcode17

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	in := "23"
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	got := letterCombinations(in)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
