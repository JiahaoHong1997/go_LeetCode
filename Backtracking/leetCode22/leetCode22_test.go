package leetcode22

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	in := 3
	want := []string{"((()))","(()())","(())()","()(())","()()()"}
	got := generateParenthesis(in)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func benchmarkDFS(b *testing.B, n int) {
	for i:=0; i<b.N; i++ {
		generateParenthesis(n)
	}
}

func benchmarkBFS(b *testing.B, n int) {
	for i:=0; i<b.N; i++ {
		generateParenthesisBFS(n)
	}
}

func BenchmarkDFS(b *testing.B) {
	benchmarkDFS(b,8)
}

func BenchmarkBFS(b *testing.B) {
	benchmarkBFS(b,8)
}