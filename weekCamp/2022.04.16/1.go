package weekCamp

import "sort"

func findClosestNumber(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    if nums[0] >= 0 {
        return nums[0]
    } else if nums[n-1] <= 0 {
        return nums[n-1]
    }
    
    a := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] < 0 {
            a = nums[i]
        } else if nums[i] == 0 {
            return 0
        } else {
            if 0-a < nums[i] {
                return a
            } else {
                return nums[i]
            }
        }
    }
    
    return 0
}