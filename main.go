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
}
