package main

import "fmt"

func main() {
	// if
	if x := 30; x == 30 {
		fmt.Println(x)
	}
	// dif ways
	if true {
		fmt.Println(true)
	}
	if !false {
		fmt.Println("!false")
	}
	// if - else if - else
	var year int = 2018
	if year == 2015 {
		fmt.Println("It's 2015")
	} else if year == 2016 {
		fmt.Println("It's 2016")
	} else if year == 2017 {
		fmt.Println("It's 2017")
	} else if year == 2018 {
		fmt.Println("It's 2018")
	} else if year == 2019 {
		fmt.Println("It's 2019")
	} else {
		fmt.Println("this year is not between 2015 and 2019")
	}
	// switch statement -- wow
	switch {
	case true:
		fmt.Println("switch")
	}

	// simple expression
	switch a := 15; {
	case true:
		fmt.Println(a)
	}

	// switch expressions
	switch a := 0; {
	case a == 1:
		fmt.Println("hello world")
	}
	//missing switch expression means "true"
	switch {
	case 200 == (10 - 5):
		fmt.Println("5 == 5")
		fallthrough
	case 100 < 1000:
		t := 500
		fmt.Println("100 < 1000", t)
		l := 1000
		fmt.Println("second statement", l)
	default:
		fmt.Println("default clause")
	case true == true:
		fmt.Println("true == true")
	}
	// the same constant
	const c = 100
	switch {
	case c > 5:
		fmt.Println("c < 5")
	case c > 75:
		fmt.Println("c < 75")
	}

	// types switch
	do(10)
	do("DiosyFútbol")
	do(true || false)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
