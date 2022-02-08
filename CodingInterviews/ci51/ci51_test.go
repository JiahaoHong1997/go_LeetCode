package ci51

import "testing"

func TestReversePairs(t *testing.T) {
	nums := []int{7,5,6,4}
	want := 5
	got := reversePairs(nums)

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}