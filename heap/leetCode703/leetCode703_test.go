package leetCode703

import (
	"reflect"
	"testing"
)

type data struct {
	in 	[]string
	out	[]int
	addNum	[]int
	k	int
	initList	[]int
}

func TestFib(t *testing.T) {
	
	d := []data {
		{
			in : []string{"KthLargest", "add", "add", "add", "add", "add"},
			out: []int{4, 5, 5, 8, 8},
			addNum : []int{0,3,5,10,9,4},
			k: 3,
			initList: []int{4, 5, 8, 2},
		},	
		{
			in: []string{"KthLargest","add","add","add","add","add"},
			out:  []int{-1,0,0,0,1},
			addNum: []int{0,-1,1,-2,-4,3},
			k: 2,
			initList: []int{0},
		},
	}

	var got [][]int
	var kl KthLargest
	for k, v := range d {
		got = append(got,[]int{})
		for i:=0; i<len(v.in); i++ {
			if v.in[i] == "KthLargest" {
				kl = Constructor(v.k, v.initList)
			} else {
				x := kl.Add(v.addNum[i])
				got[k] = append(got[k],x)
			}
		}
	}

	want := [][]int{}
	for i:=0; i<len(d); i++ {
		want = append(want,d[i].out)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v",want,got)
	}
}