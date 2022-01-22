package main

import "fmt"

func func1() func(int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}

func func2() (val int) {
	val = 10
	defer func() {
		val += 1
	}()
	return val
}

func func3() int {
	val := 10
	defer func() {
		val += 1
	}()
	return val
}

func printFunc() {
	sumFunc := func1()
	fmt.Println(sumFunc(1)) // 1
	fmt.Println(sumFunc(4)) // 5

	fmt.Println(func2()) // 11

	fmt.Println(func3()) // 10
}

func main() {
	printFunc()

	a := 10
	defer func() {
		// a = 20
		fmt.Println("defer: a=", a)
	}()
	
	fmt.Println("origin: a=", a)
	a = 30 
	fmt.Println("revise a=", a)

	s := "\""
	fmt.Println(s)

	b := byte(2+'0')
	fmt.Printf("b:%s\n",string(b))

	var c float64
	var d float64
	c = 1 - (2*1.0) / 3
	d = 4 / c
	e := 12.0
	isEqual := e == d
	fmt.Println(c,d,isEqual)
}
