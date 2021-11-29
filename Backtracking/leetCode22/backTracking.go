package leetcode22

func generateParenthesis(n int) []string {
    ret := []string{}
    path := []byte{}
    backTracking(n, path, &ret, 0, 0)
    return ret
}

func backTracking(n int, path []byte, ret *[]string, count1 int, count2 int) {
    if len(path) == 2*n {
        dst := make([]byte,2*n)
        copy(dst,path)
        *ret = append(*ret, string(dst))
        return
    }

    if count1 < n {
        path = append(path, '(')
        backTracking(n, path, ret, count1+1, count2)
        path = path[:len(path)-1]
    }

    if count1 > count2 {
        path = append(path, ')')
        backTracking(n, path, ret, count1, count2+1)
        path = path[:len(path)-1]
    }
}