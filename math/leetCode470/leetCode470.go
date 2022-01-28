package leetcode470

import "math/rand"

func rand10() int {
    x := rand7()
    for x==7 || x==6 {
        x = rand7()
    } 
    
	y := rand7()
	for y==7 {
        y = rand7()
	}

	if y>=1 && y<=3 {
		return x
	}
	return x+5
}

func rand7() int {
	return rand.Intn(7)+1
}