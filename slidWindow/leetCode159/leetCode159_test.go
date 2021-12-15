package leetcode159

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	s := []string{
		"eceba",
		"ccaabbb",
		"asdbsjdnajdnaskjdbejhkbbbdbewjkldbqwiwquibdakbcwjhdsdnasjkdnvgfwuycbewk",
		"jjjdnasjbbbbbbskakkkshcnwbiucewkjbnfweijvbknfoewiufbnbfnfhdsndnjfnfjfjhdenjednhjcjkllwqbqwomqpj",
	}

	want := []int{3,5,5,7}
	got := []int{}

	for i:=0; i<len(s); i++ {
		got = append(got, lengthOfLongestSubstringTwoDistinct(s[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want,got)
	}
}