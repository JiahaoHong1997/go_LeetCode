package ci46

func translateNum(num int) int {
	if num == 0 {
		return 1
	}

	arr := []int{}
	for num > 0 {
		arr = append([]int{num%10}, arr...)
		num /= 10 
	}

	ret := 0
	var backTracking func([]int)
	backTracking = func(nums []int) {
		n := len(nums)
		if n == 0 {
			ret++
			return 
		}

		
		backTracking(nums[1:])
		if n >= 2 && nums[0] != 0{ 
			x, y := nums[0], nums[1]
			t := x*10+y
			if t <= 25 {
				backTracking(nums[2:])
			}
		}
	}

	backTracking(arr)
	return ret
}