package leetcode394

import (
	"reflect"
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := []string{
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
		"abc3[cd]xyz",
		"3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
	}

	want := []string{
		"aaabcbc",
		"accaccacc",
		"abcabccdcdcdef",
		"abccdcdcdxyz",
		"zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
	}

	got := []string{}

	for i:=0; i<len(s); i++  {
		got = append(got, decodeString(s[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
