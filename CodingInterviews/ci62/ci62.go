package ci62

func lastRemaining(n int, m int) int {

	ret := 0
	for i:=2; i<=n; i++ {
		ret = (ret+m)%i
	}
	return ret
}