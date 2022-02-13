package ci61
import "sort"

func isStraight(nums []int) bool {

	sort.Slice(nums, func(i,j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	var zeroCount func() int
	zeroCount = func() int {
		x := 0
		for i:=0; i<5; i++ {
			if nums[i] == 0 {
				x++
			}
		}
		return x
	}

	z := zeroCount()
	if z == 0 {
		x := nums[0]
		for i:=1; i<5; i++ {
			if nums[i] != x+1 {
				return false
			}
			x = nums[i]
		}
		return true
	}

    var x int
	for i:=0; i<5; i++ {
        if nums[i] == 0 {
            continue
        }
        if x == 0 && nums[i] != 0 {
            x = nums[i]
            continue
        }
        if nums[i] == x {
            return false
        }
		if nums[i] != x+1 {
			z -= (nums[i]-x-1)
			if z < 0 {
				return false
			} 
		}
		x = nums[i]
	}
	return true
}