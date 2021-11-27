package leetcode78

func subsets(nums []int) [][]int {
    path := make([]int,0)
    ret := make([][]int,0)
    DFS(nums,&ret,path,0)
    return ret
}

func DFS(nums []int, ret *[][]int, path []int, start int) {
    dst := make([]int,len(path))
    copy(dst,path)
    *ret = append(*ret, dst)

    for i:=start; i<len(nums); i++ {
        path = append(path, nums[i])
        DFS(nums,ret,path,i+1)
        path = path[:len(path)-1]
    }
}