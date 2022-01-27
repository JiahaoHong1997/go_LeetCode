package ci10

func numWays(n int) int {
    mod := 1000000007
    if n == 0 || n == 1 {
        return 1
    }

    first := 1
    second := 1
    ret := 0
    for i:=2; i<=n; i++ {
        ret = (first + second) % mod
        first = second
        second = ret
    }
    return ret
}