package leetcode583

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1 := "abcde"
	word2 := "ace"

	got := minDistance(word1,word2)
	want := 2


	if want != got {
		t.Errorf("expected%v, got%v", want, got)
	}
}