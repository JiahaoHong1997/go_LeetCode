package leetcode76

import (
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := []string{
		"ADOBECODEBANC",
		"a",
		"a",
	}

	tmp := []string{
		"ABC",
		"a",
		"aa",
	}

	want := []string{
		"BANC",
		"a",
		"",
	}

	got := []string{}

	for i:=0; i<len(s); i++ {
		got = append(got, minWindow(s[i],tmp[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected:%v, got:%v", want, got)
	}
}