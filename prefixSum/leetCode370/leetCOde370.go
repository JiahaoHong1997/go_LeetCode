package leetcode370

func getModifiedArray(length int, updates [][]int) []int {

	nums := make([]int, length)
	diff := make([]int, length)

	for i:=0; i<len(updates); i++ {
		left, right, val := updates[i][0], updates[i][1], updates[i][2]
		diff[left] += val
		if right+1 >= length {
			continue
		}
		diff[right+1] -= val
	}

	nums[0] = diff[0]
	for i:=1; i<length; i++ {
		nums[i] = nums[i-1]+diff[i]
	}
	
	return nums
}