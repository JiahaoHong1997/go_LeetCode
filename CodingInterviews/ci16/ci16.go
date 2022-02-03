package ci16

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}

	t := myPow(x, n/2)
	if n % 2 == 0 {
		return t * t
	}
	return t * t * x
}