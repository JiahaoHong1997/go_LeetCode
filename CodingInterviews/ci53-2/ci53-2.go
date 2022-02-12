package ci53

func missingNumber(nums []int) int {
	var sum int
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}

	n := len(nums)
	return (1+n)*n/2 - sum
}