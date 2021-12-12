package leetcode1060

func missingElement(nums []int, k int) int {
	n := len(nums)
	if nums[n-1]-nums[0] < k+(n-1) {
		return nums[n-1] + (k + (n - 1) - (nums[n-1] - nums[0]))
	}

	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid]-nums[0] == k+mid {
			r = mid
		} else if nums[mid]-nums[0] < k+mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[0] + k + (l - 1)
}
