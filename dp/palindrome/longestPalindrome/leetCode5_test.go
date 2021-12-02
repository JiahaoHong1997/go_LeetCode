package longestpalindrome

import (
	"reflect"
	"testing"
)

type leetcode5Test struct {
	in   []string
	got  []string
	want []string
}

func TestLongestpalindromeDP(t *testing.T) {
	l := leetcode5Test{
		in: []string{
			"babad",
			"cbbd",
			"a",
			"ac",
		},
		want: []string{
			"bab",
			"bb",
			"a",
			"a",
		},
	}

	for i:=0; i<len(l.in); i++ {
		l.got = append(l.got, longestpalindromeDP(l.in[i]))
	}

	if !reflect.DeepEqual(l.got, l.want) {
		t.Errorf("expected:%v, got:%v", l.want, l.got)
	}
}
