package leetcode39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
    ret := make([][]int,0)
    l := make([]int,0)
    sort.Slice(candidates,func(i,j int)bool{
        return candidates[i] < candidates[j]
    })

    var backTracking func(int, int)
    backTracking = func(k int, start int) {
        if k == 0 {
            l1 := make([]int,len(l))
            copy(l1,l)
            ret = append(ret, l1)
            return
        }
        if k < candidates[0] {
            return
        }
        

        for i:=start; i<len(candidates); i++ {
            if k >= candidates[i] {
                l = append(l,candidates[i])
                backTracking(k-candidates[i],i)  // // 从当前节点往后寻找，避免出现重复
                l = l[:len(l)-1]
            }
        }
    }
    backTracking(target,0)
    return ret
}