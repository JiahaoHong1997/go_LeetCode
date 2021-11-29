package leetcode79

import (
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	got := exist(board,word)
	want := true

	if got != want {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}