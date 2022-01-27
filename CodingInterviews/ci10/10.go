package ci10

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } 

    mod := 1000000007
    first := 0
    second := 1
    ret := 1
    for i:=2; i<=n; i++ {
        ret = int((first + second) % mod)
        first = second
        second = ret
    }
    return ret
}