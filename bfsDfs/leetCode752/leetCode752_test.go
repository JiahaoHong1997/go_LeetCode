package leetcode752

import (
	"testing"
	"reflect"
)

func TestOpenLock(t *testing.T) {
	deadens := [][]string{
		{"0201","0101","0102","1212","2002"},
		{"0201","0101","0102","1212","2002"},
		{"8888"},
		{"8887","8889","8878","8898","8788","8988","7888","9888"},
		{"0000"},
	}

	target := []string{
		"0202",
		"0000",
		"0009",
		"8888",
		"8888",
	}

	want := []int{
		6,
		0,
		1,
		-1,
		-1,
	}

	got := []int{}

	for i:=0; i<len(deadens); i++ {
		got = append(got, openLock(deadens[i],target[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
} 