package leetcode17

func letterCombinations(digits string) []string {
    if digits == "" {
        return nil
    }
    ret := []string{}
    path := []byte{}
    numToWord := [][]byte{
        {},
        {},
        {'a','b','c'},
        {'d','e','f'},
        {'g','h','i'},
        {'j','k','l'},
        {'m','n','o'},
        {'p','q','r','s'},
        {'t','u','v'},
        {'w','x','y','z'},
    }
    backTracking([]byte(digits), path, &ret, numToWord, 0)
    return ret
}

func backTracking(digits []byte, path []byte, ret *[]string, wordList [][]byte, start int) {
    if len(digits) == 0 {
        dst := make([]byte,len(path))
        copy(dst,path)
        *ret = append(*ret, string(dst))
        return
    }

    for i:=start; i<len(digits); i++ {
        num := digits[i]
        for _,v := range wordList[int(num)-'0'] {
            path = append(path, v)
            digits = append(digits[:i], digits[i+1:]...)
            backTracking(digits, path, ret, wordList, i)
            digits = append(digits[:i], append(append([]byte{num}, digits[i:]...))...)
            path = path[:len(path)-1]
        }   
    }
}