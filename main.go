package main

import (
	"context"
	"fmt"
	"time"
)

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

	f := []int{1,2,3,4}
	g := f[:0]
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))
	g = append(g, 5, 6, 7)
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))
	f = g
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))

	// ch1 := make(chan int,1)
	// ch2 := make(chan int,1)
	// i, j := 1, 2

	// f1 := func() {
	// 	if i < 100 {
	// 		fmt.Println(i)
	// 		i += 2
	// 		ch2<-1	
	// 	}
	// }
	
	// f2 := func() {
	// 	if j <= 100 {
	// 		fmt.Println(j)
	// 		j += 2
	// 		ch1<-1
	// 	}
	// }
	
	// ch1<-1
	// for {
	// 	select {
	// 	case <-ch1:
	// 		f1()
	// 	case <-ch2:
	// 		f2()
	// 	default:
	// 		return
	// 	}
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(c context.Context) {

		for {
			select {
			case <-ctx.Done():
				fmt.Println("processA is over!")
			default:
				time.Sleep(1*time.Second)
				fmt.Println("AAA!")
			}
		}
	}(ctx)

	go func(c context.Context) {

		for {
			select {
			case <-ctx.Done():
				fmt.Println("processB is over!")
			default:
				time.Sleep(1*time.Second)
				fmt.Println("BBB!")
			}
		}
	}(ctx)
	
	select {
	case <-ctx.Done():
		return
	}
}