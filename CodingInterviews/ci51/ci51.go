package ci51


func reversePairs(nums []int) int {
	if len(nums) == 0 {
        return 0
    }
	count := 0
	
	var mergeSort func([]int) []int
	mergeSort = func(nums []int) []int {

		if len(nums) == 1 {
			return nums
		}

		n := len(nums)
		a := mergeSort(nums[:n/2])
		b := mergeSort(nums[n/2:])
		
		tmp := make([]int, 0)
		i, j := 0, 0
		for i < len(a) || j < len(b) {
			if i == len(a) {
				tmp = append(tmp, b[j:]...)
				break
			} else if j == len(b) {
				tmp = append(tmp, a[i:]...)
				break
			}

			if a[i] > b[j] {
				tmp = append(tmp, b[j])
				j++
				count += len(a) - i
			} else {
				tmp = append(tmp, a[i])
				i++
			}
		}
		return tmp
	}
	mergeSort(nums)
	return count
}