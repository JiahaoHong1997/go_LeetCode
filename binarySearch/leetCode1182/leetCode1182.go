package leetcode1182

func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(queries)
	ret := make([]int, 0)

	set1 := []int{}
	set2 := []int{}
	set3 := []int{}

	for i:=0; i<len(colors); i++ {
		if colors[i] == 1 {
			set1 = append(set1, i)
		} else if colors[i] == 2 {
			set2 = append(set2, i)
		} else {
			set3 = append(set3, i)
		}
	}

	for i:=0; i<n; i++ {
		if colors[queries[i][0]] == queries[i][1] {
			ret = append(ret, 0)
			continue
		}

		var set []int
		if queries[i][1] == 1 {
			set = set1
		} else if queries[i][1] == 2 {
			set = set2
		} else {
			set = set3
		}

		if len(set) == 0 {
			ret = append(ret, -1)
		} else {
			ret = append(ret, search(set, queries[i][0]))
		}
		
	}
	return ret
}

func search(nums []int, target int) int {
	n := len(nums)
	if target > nums[n-1] {
		return target - nums[n-1]
	}
	if target < nums[0] {
		return nums[0] - target
	}

	l, r := 0, n-1
	for l < r {
		mid := l+(r-l)/2
		if nums[mid] == target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return min(target-nums[l-1], nums[l]-target)
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}