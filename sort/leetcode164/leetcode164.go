package leetcode164

func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }

    var maxNum func(...int) int
    maxNum = func(a ...int) int {
        max := a[0]

        for i:=1; i<len(a); i++ {
            if a[i] > max {
                max = a[i]
            }
        }

        return max
    }

    n := len(nums)
    for exp:=1; exp<=maxNum(nums...); exp *= 10 {
        buf := make([]int, n)
        cnt := [10]int{}
        for i:=0; i<n; i++ {
            bit := nums[i]/exp%10
            cnt[bit]++
        }

        for i:=1; i<10; i++ {
            cnt[i] += cnt[i-1]
        }

        for i:=n-1; i>=0; i-- {
            bit := nums[i]/exp%10
            buf[cnt[bit]-1] = nums[i]
            cnt[bit]--
        }

        copy(nums, buf)
    } 

    ret := nums[1]-nums[0]
    for i:=2; i<n; i++ {
        if nums[i]-nums[i-1] > ret {
            ret = nums[i]-nums[i-1]
        }
    }

    return ret
}
