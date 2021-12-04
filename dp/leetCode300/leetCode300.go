package leetcode300

func lengthOfLIS(nums []int) int {
	cell := make([]int, 1)
	cell[0] = nums[0]

	for i:=1; i<len(nums); i++ {
		if nums[i] > cell[len(cell)-1] {
			cell = append(cell, nums[i])
		} else {
			tmp := findLocation(cell, nums[i])
			if tmp == -1 {
				continue
			} else {
				cell[tmp] = nums[i]
			}
		}
	}

	return len(cell)
}

func findLocation(cell []int, target int) int {
	n := len(cell)
	if target <= cell[0] {
		return 0
	}
    if  target == cell[n-1] {
        return -1
    }
	if len(cell) >= 2 && target > cell[n-2] {
		return n-1
	} 

	l, r := 0, n-1
	for l < r {
		mid := l+(r-l)/2
		if cell[mid] == target {
			l = mid + 1
		} else if cell[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if cell[l-1] == target {
		return -1
	} 
	return l
}