package leetcode249

func groupStrings(strings []string) [][]string {
    m := make(map[string]bool)
    traveled := make(map[string]int)
    for i:=0; i<len(strings); i++ {
        m[strings[i]] = true
        traveled[strings[i]]++
    }

    var judgeUp func(string) string 
    judgeUp = func(s string) string {

        var ret string
        for i:=0; i<len(s); i++ {
            x := s[i]
            if x == 'a' {
                ret += "z"
            } else {
                ret += string(x-1)
            }
        }
        return ret
    } 

    var judgeDown func(string) string
    judgeDown = func(s string) string {

        var ret string
        for i := 0; i<len(s); i++ {
            x := s[i]
            if x == 'z' {
                ret += "a"
            } else {
                ret += string(x+1)
            }
        }

        return ret
    }

    ret := [][]string{}
    for i:=0; i<len(strings); i++ {
        if traveled[strings[i]] == 0 {
            continue
        }

        t := []string{}
        x := strings[i]

        for x[0] != 'a' {
            if m[x] {
                for traveled[x] > 0 {
                    t = append([]string{x}, t...)
                    traveled[x]--
                } 
            }
            
            x = judgeUp(x)
        }

        for m[x] && traveled[x] > 0 {
            t = append([]string{x}, t...)
            traveled[x]--
        }

        x = judgeDown(strings[i])
        for x[0] != 'z' {
            if m[x] {
                for traveled[x] > 0 {
                    t = append(t, x)
                    traveled[x]--
                }
            }
            
            x = judgeDown(x)
        }

        for x[0] == 'z' && m[x] && traveled[x] > 0 {
            t = append(t, x)
            traveled[x]--
        }

        ret = append(ret, t)
    }

    return ret
}