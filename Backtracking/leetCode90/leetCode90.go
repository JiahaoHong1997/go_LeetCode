package leetCode90

import "sort"

func subsetsWithDup(nums []int) [][]int {
    ret := make([][]int,0)
    path := make([]int,0)
	sort.Ints(nums)
    DFS(nums,path,&ret,0)

    return ret
}

func DFS(nums []int, path []int, ret *[][]int, start int) {
    dst := make([]int,len(path))
    copy(dst,path)
    *ret = append(*ret,dst)

    for i:=start; i<len(nums); i++ {
        if i>start && nums[i]==nums[i-1] {
            continue
        }
        path = append(path,nums[i])
        DFS(nums,path,ret,i+1)
        path = path[:len(path)-1]
    }
}