package leetcode1318

import (
	"reflect"
	"testing"
)

func TestMinFlips(t *testing.T) {
	a := []int{2,4,1,432,543,6341,36242,734253}
	b := []int{6,2,2,526,362,426322,432342,436532}
	c := []int{5,7,3,5432,745634,43242,54362,73723532}

	want := []int{3,1,0,6,14,10,10,16}
	got := []int{}

	for i:=0; i<len(a); i++ {
		got = append(got, minFlips(a[i],b[i],c[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v",want, got)
	}
	
}