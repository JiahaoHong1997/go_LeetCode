package leetcode46

func permute(nums []int) [][]int {
    ret := make([][]int,0)
    path := make([]int,0)
    n := len(nums)
    backTrace(nums, n, path, &ret)
    return ret
}

func backTrace(nums []int, length int, path []int, ret *[][]int) {
    if length == 0 {
        p := make([]int,len(path))
        copy(p,path)
        *ret = append(*ret,p)
        return
    }

    for i:=0; i<length; i++ {
        x := nums[i]
        path = append(path,x)
        nums = append(nums[:i],nums[i+1:]...)
        backTrace(nums, len(nums), path, ret)
        nums = append(nums[:i],append([]int{x},nums[i:]...)...)
        path = path[:len(path)-1]
    }
}