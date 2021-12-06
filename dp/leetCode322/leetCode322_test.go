package leetcode322

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := [][]int{{1,2,5}}
	amount := []int{100}

	want := []int{20}
	got := []int{}
	for i:=0; i<len(coins); i++ {
		got = append(got, coinChange(coins[i],amount[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}