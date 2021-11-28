package leetcode40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    path := make([]int,0)
    ret := make([][]int,0)
    sort.Ints(candidates)
    backTracking(0,candidates,target,path,&ret)
    return ret
}

func backTracking(start int, candidates []int, target int, path []int, ret *[][]int) {
    if target == 0 {
        dst := make([]int,len(path))
        copy(dst,path)
        *ret = append(*ret,dst)
        return
    } else if target < 0 {
        return
    } 

    for i:=start; i<len(candidates); i++ {
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }
        x := candidates[i]
        path = append(path, x)
        backTracking(i+1,candidates,target-x,path,ret)
        path = path[:len(path)-1]
    }
}