package ci33

func verifyPostorder(postorder []int) bool {

    if len(postorder) == 0 {
		return true
	}

	n := len(postorder)
	x := postorder[n-1]

	t := -1
	for i:=0; i<n-1; i++ {
		if postorder[i] > x {
			t = i
			break
		}
	}

    if t != -1 {
       for i:=t+1; i<n-1; i++ {
		    if postorder[i] < x {
			    return false
		    }
	    } 
    }

	return verifyPostorder(postorder[:n-1])
}