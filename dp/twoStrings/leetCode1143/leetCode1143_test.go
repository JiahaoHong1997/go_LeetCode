package leetcode1143

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"

	got := longestCommonSubsequence(text1,text2)
	want := 3


	if want != got {
		t.Errorf("expected%v, got%v", want, got)
	}
}