package leetcode443

func compress(chars []byte) int {
    if len(chars) == 1 {
        return 1
    }
    count := 1
    start := 0
    c := chars[0]

    ret := 0
    for i:=1; i<len(chars); i++ {
        if chars[i] == c && i != len(chars)-1 {
            count++
            continue
        } else {
            var end int
            if i == len(chars)-1 && chars[i] == c {
                count++
                end = i+1
            } else if i == len(chars)-1{
                ret++
                end = i
            } else {
                end = i
            }
            var b []byte
            if count <= 2 {
                ret += count
                if count == 2 {
                    if i == len(chars)-1 && chars[i] == c {
                        chars[i] = '2'
                    } else {
                        chars[i-1] = '2'
                    }   
                }
            } else {
                t := count
                chars = append(chars[:start+1], chars[end:]...)
                for t > 0 {
                    b = append([]byte{byte(t%10+'0')}, b...)
                    t /= 10
                    ret++
                }
                chars = append(chars[:start+1], append(b, chars[start+1:]...)...)
                
                ret++
                i = len(b)+start+1
            }

            if i < len(chars) {
                c = chars[i]
            }
            count = 1
            start = i
        }
    }

    return ret
}