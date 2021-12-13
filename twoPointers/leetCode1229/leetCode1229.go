package leetcode1229

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	n1, n2 := len(slots1), len(slots2)
	i, j := 0, 0
	for i < n1 || j < n2 {
		if i == n1 {
			if slots2[j][1] <= slots1[i-1][0] {
				j++
				continue
			} else if slots2[j][0] >= slots1[i-1][1] {
				break
			} else {
				if min(slots1[i-1][1],slots2[j][1]) - max(slots1[i-1][0],slots2[j][0]) >= duration {
					return []int{max(slots1[i-1][0],slots2[j][0]),max(slots1[i-1][0],slots2[j][0])+duration}
				} else {
					j++
					continue
				}
			}
		}

		if j == n2 {
			if slots1[i][1] <= slots2[j-1][0] {
				i++
				continue
			} else if slots1[i][0] >= slots2[j-1][1] {
				break
			} else {
				if min(slots1[i][1],slots2[j-1][1]) - max(slots1[i][0],slots2[j-1][0]) >= duration {
					return []int{max(slots1[i][0],slots2[j-1][0]),max(slots1[i][0],slots2[j-1][0])+duration}
				} else {
					i++
					continue
				}
			}
		}
		if slots1[i][1] <= slots2[j][0] {
			i++
		} else if slots2[j][1] <= slots1[i][0] {
			j++
		} else {
			if min(slots1[i][1],slots2[j][1]) - max(slots1[i][0],slots2[j][0]) >= duration {
				return []int{max(slots1[i][0],slots2[j][0]),max(slots1[i][0],slots2[j][0])+duration}
			} else if slots1[i][1] < slots2[j][1] {
				i++
			} else if slots1[i][1] > slots2[j][1] {
				j++
			} else {
				i++
				j++
			}
		}
	}
	return []int{}
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}