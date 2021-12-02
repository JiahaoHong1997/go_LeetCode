package longestPalindromeSequence

import (
	"reflect"
	"testing"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	in := []string{
		"abc",
		"aaa",
		"babad",
		"cbbd",
		"shduvbsdfuudsxneuuenx",
		"bbbab",
		"cddb",
	}

	want := []int{1,3,3,2,8,4,2}
	got := []int{}

	for i:=0; i<len(in); i++ {
		got = append(got, longestPalindromeSubseq(in[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}