package leetcode267

func generatePalindromes(s string) []string {

    cnt := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        cnt[s[i]]++
    }

    letterSet := make([]byte, 0)
    ret := []string{}
    oddNum := 0
    var midLetter byte
    for k, v := range cnt {
        if v%2 == 1 {
            oddNum++
            midLetter = k
            if oddNum > 1 {
                return ret
            }
        }

        n := v/2
        for j:=0; j<n; j++ {
            letterSet = append(letterSet, k)
        }
    }

    path := []byte{}
    allPath := [][]byte{}
    var backTracking func([]byte)
    backTracking = func(ss []byte) {
        if len(ss) == 0 {
            t := make([]byte, len(path))
            copy(t, path)
            allPath = append(allPath, t)
            return
        }
        m := make(map[byte]bool)

        for i:=0; i<len(ss); i++ {
            x := ss[i]
            if m[x] {
                continue
            }

            m[x] = true
            ss = append(ss[:i], ss[i+1:]...)
            path = append(path, x)
            backTracking(ss)
            path = path[:len(path)-1]
            ss = append(ss[:i], append([]byte{x}, ss[i:]...)...)
        }
    }

    var reverse func([]byte) []byte
    reverse = func(b []byte) []byte {
        b1 := make([]byte, len(b))
        for i:=0; i<len(b); i++ {
            b1[len(b)-1-i] = b[i]
        }

        return b1
    }

    backTracking(letterSet)

    for i:=0; i<len(allPath); i++ {
        b := allPath[i]
        b1 := reverse(b)
        if oddNum == 0 {
            ret = append(ret, string(b)+string(b1))
        } else {
            ret = append(ret, string(b)+string(midLetter)+string(b1))
        }
    }

    return ret
}