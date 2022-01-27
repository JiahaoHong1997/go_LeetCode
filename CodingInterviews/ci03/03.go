package ci03

import "sort"

func findRepeatNumber(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        if nums[i] < nums[j] {
            return true
        }
        return false
    })

    x := nums[0]
    for i:=1; i<len(nums); i++ {
        if nums[i] == x {
            return x
        }
        x = nums[i]
    }
    return 0
}