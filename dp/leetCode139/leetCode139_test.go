package leetcode139

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	str := []string{
		"leetcode",
		"applepenapple",
		"catsandog",
	}

	wordDicts := [][]string{
		{"leet", "code"},
		{"apple", "pen"},
		{"cats", "dog", "sand", "and", "cat"},
	}

	var got []bool
	for i:=0; i<len(str); i++ {
		got = append(got, wordBreak(str[i],wordDicts[i]))
	}

	want := []bool{true,true,false}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

