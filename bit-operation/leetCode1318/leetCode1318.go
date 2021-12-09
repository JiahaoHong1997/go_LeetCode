package leetcode1318

func minFlips(a int, b int, c int) int {
	count := 0
    for a != 0 || b != 0 || c != 0 {
        var bitC int
        if c^(c-1) == 1 {
            bitC = 1
        }

        var bitA int
        var bitB int
        if a^(a-1) == 1 {
            bitA = 1
        }
        if b^(b-1) == 1 {
            bitB = 1
        }

        if bitC == 1 && (bitA | bitB) == 0 {
            count++
        } else if bitC == 0 && (bitA | bitB) == 1 {
            if bitA & bitB == 1 {
                count += 2
            } else {
                count++
            }
        }

        a >>= 1
        b >>= 1
        c >>= 1
    }

    return count
}