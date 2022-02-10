package ci53

func search(nums []int, target int) int {

    if len(nums) == 0 {
        return 0
    }
	var findLeftBoundry func() int 
	findLeftBoundry = func() int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] == target {
				r = mid
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
		
	var findRightBoundry func() int 
	findRightBoundry = func() int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] == target {
				l = mid+1
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}


	if target < nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return 0
	}

	left := findLeftBoundry()
	if nums[left] != target {
		return 0
	}
	right := findRightBoundry()

    if nums[right] == target {
        return right - left + 1
    }
	return right-left
}