package leetcode29

import (
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {
	dividend := []int{10,7,-2147483648,-2147483648}
	divisor := []int{3,-3,1,-1}
	want := []int{3,-2,-2147483648,2147483647}

	got := []int{}
	for i:=0; i<len(dividend); i++ {
		got = append(got, divide(dividend[i],divisor[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want,got)
	}
}