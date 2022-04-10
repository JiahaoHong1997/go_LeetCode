package weekCamp

import "sort"

func maximumProduct(nums []int, k int) int {
	const Max = 1000000007

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	i := len(nums)-1
	x := nums[i]
	for k > 0 {
		if i >= 0 {
			if nums[i] == x {
				nums[i]++
				i--
				k--
			} else if nums[i] >= x+1 {
                i = len(nums)-1
                x = nums[i]
			} 
		} else {
			i = len(nums)-1
			x = nums[i]
		}
        //fmt.Println(nums)
	}
    
    //fmt.Println(nums)
	ret := 1
	for i:=0; i<len(nums); i++ {
		ret = ret*nums[i]%Max
	}

	return ret
}