package ci31

func validateStackSequences(pushed []int, popped []int) bool {
    if len(pushed) == 0 {
        return true
    }
    m := make(map[int]bool, 0)
	arr := make([]int, 0)
	x := popped[0]
	t := 0
	for i:=0; i<len(pushed); i++ {
		m[pushed[i]] = true
		arr = append(arr, pushed[i])
		if x == pushed[i] {
			for m[x] {
                m[x] = false
				if len(arr) > 0 && arr[len(arr)-1] != x {
					return false
				} else {
					t++
                    if t < len(popped) {
                        x = popped[t]
                    }
                    if len(arr) > 0 {
                        arr = arr[:len(arr)-1]
                    }
				}
                
 			}
		}
	}

	return len(arr) == 0
}