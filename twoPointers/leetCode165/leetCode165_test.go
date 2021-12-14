package leetcode165

import (
	"reflect"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	version1 := []string{
		"1.01",
		"1.0",
		"0.1",
		"1.0.1",
		"7.5.2.4",
	}

	version2 := []string{
		"1.001",
		"1.0.0",
		"1.1",
		"1",
		"7.5.3",
	}

	want := []int{0,0,-1,1,-1}
	got := []int{}

	for i:=0; i<len(version1); i++ {
		got = append(got, compareVersion(version1[i],version2[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}