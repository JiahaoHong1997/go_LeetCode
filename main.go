package main

import (
	"context"
	"encoding/json"
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

func String(v string) *string { return &v }

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

	b := byte(2 + '0')
	fmt.Printf("b:%s\n", string(b))

	var c float64
	var d float64
	c = 1 - (2*1.0)/3
	d = 4 / c
	e := 12.0
	isEqual := e == d
	fmt.Println(c, d, isEqual)

	f := []int{1, 2, 3, 4}
	g := f[:0]
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))
	g = append(g, 5, 6, 7)
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))
	f = g
	fmt.Printf("f: %p, %v, %v, %v\n", f, f, len(f), cap(f))
	fmt.Printf("g: %p, %v, %v, %v\n", g, g, len(g), cap(g))

	// var sourceCode *string
	// sourceCode = nil
	// *sourceCode = ""
	// fmt.Println("sourceCode:", *sourceCode)
	now := time.Now()
	xxx := 2.5
	twoHoursAgo := now.Add(time.Duration(-1*xxx*60)*time.Minute)
	fmt.Println("now:", now, "\ntwoHours:", twoHoursAgo)

	addTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	addTime = time.Unix(now.Unix(), 0)
	fmt.Println("zeroDayTime:", addTime)

	value := "{\"warehouse1\":1,\"warehouse2\":0.5}"
	valueMap := make(map[string]float64)
	err := json.Unmarshal([]byte(value), &valueMap)
	if err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Println("success:", valueMap["warehouse1"], valueMap["warehouse2"])
	}

	aList := make([]string, 0)
	bList := []string{"1","2","3","4"}
	codeList := make([]string, 0)
	copy(codeList, aList)
	codeList = append(codeList, bList...)
	fmt.Println("codeList", codeList)

	type Example struct {
		Type  *string
		Name  *string
		Phone *string
	}

	cType := []string{"manager", "receiver", "reviseReceiver"}
	name := []string{"小红", "小蓝", "小驴"}
	phone := []string{"2141432", "4235243", "5312342523"}
	contactList := []*Example{
		&Example{
			Type:  &cType[0],
			Name:  &name[0],
			Phone: &phone[0],
		},
		&Example{
			Type:  &cType[1],
			Name:  &name[1],
			Phone: &phone[1],
		},
		&Example{
			Type:  &cType[2],
			Name:  &name[2],
			Phone: &phone[2],
		},
	}

	contactJson, _ := json.Marshal(contactList)
	fmt.Println("1231241", string(contactJson))

	// contactList1 := []*Example{}
	ttestt := make(map[string]interface{}, 0)
	testJson := `{"min":0,"max":100,"samplingRatio":1}`
	err = json.Unmarshal([]byte(testJson), &ttestt)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	fmt.Println("213143", ttestt["max"])
	// fmt.Println("213143", *contactList1[0].Name, *contactList1[1].Phone, *contactList1[2].Type)
	
	resMap := make(map[string][]string, 0)
	if _, ok := resMap["xiaoliu"]; !ok {
		resMap["xiaoliu"] = make([]string, 0)
	}
	resMap["xiaoliu"] = append(resMap["xiaoliu"], "xingming")
	fmt.Println(resMap)

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
				time.Sleep(1 * time.Second)
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
				time.Sleep(1 * time.Second)
				fmt.Println("BBB!")
			}
		}
	}(ctx)

	test := &Example{
		Type:  &cType[0],
		Name:  &name[0],
		Phone: &phone[0],
	}

	bbbb, _ := json.Marshal(test)
	fmt.Println(string(bbbb))

	select {
	case <-ctx.Done():
		return
	}

	
}
