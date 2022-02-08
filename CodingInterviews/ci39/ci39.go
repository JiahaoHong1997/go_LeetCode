package ci39

func majorityElement(nums []int) int {
	tmp := nums[0]
	count := 1

	for i:=1; i<len(nums); i++ {
		if count == 0 {
			tmp = nums[i]
			count = 1
			continue
		}
		if nums[i] == tmp {
			count++
		} else {
			count--
		}
	}

	return tmp
}