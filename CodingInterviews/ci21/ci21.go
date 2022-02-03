package ci21


func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	} 
	t := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] % 2 == 0 {
			t = i
			break
		}
	}

	for i:=t+1; i<len(nums); i++ {
		if nums[i] % 2 == 1 {
			nums[t], nums[i] = nums[i], nums[t]
			t++
		}
	}

	return nums
}