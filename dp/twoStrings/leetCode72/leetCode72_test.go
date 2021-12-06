package leetcode72

import (
	"reflect"
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1 := []string{
		"horse",
		"intention",
	} 

	word2 := []string{
		"ros",
		"execution",
	}

	want := []int{3,5}
	got := []int{}
	for i:=0; i<len(word1); i++ {
		got = append(got, minDistance(word1[i],word2[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}