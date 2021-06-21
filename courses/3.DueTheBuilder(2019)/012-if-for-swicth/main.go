package main

import "fmt"

func main() {
	// >>>>>>>> if
	x := 20
	// no idiomatic
	if x == 10 {
		fmt.Println("x == 10")
	} else {
		fmt.Println("x != 10")
	}
	// idiomatic
	if x == 10 {
		fmt.Println("x == 10")
		return // exit from main
	}
	fmt.Println("x != 10")

	// call
	if r, err := f(5); err == nil {
		fmt.Println(r)
		return
	} else {
		fmt.Println(err)
	}

	// >>>>>>>> swicth
	// expression
	z := 10
	switch z == 10 {
	case z == 10:
		fmt.Println("z == 10")
	case z == 20:
		fmt.Println("z == 20")
	default:
		fmt.Println("I don't know what is z")
	}

	// expre + comparation
	switch pi := 3.1415; pi == 3.1415 {
	case pi == 3.13:
		fmt.Println("pi == 3.13")
	case pi == 3.12:
		fmt.Println("pi == 3.12")
	case pi == 3.1415:
		fmt.Println("pi == 3.1415")
	default:
		fmt.Println("I don't know what is pi")
	}

	// without expression
	switch {
	case 5 > 10:
		fmt.Println("5 > 10")
	case 15 < 10:
		fmt.Println("15 > 10")
	case 3 > 1:
		fmt.Println("3 > 1") // only result
		fallthrough          // force next case
	case 3 > 2:
		fmt.Println("3 > 2")
	}

	// >>>>>>>> for
	fmt.Println(">>>>>>>>>>>>> for loop")
	// three-component loop
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum:", sum) // 10

	// while loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("n:", n) // 8 (1*2*2*2)

	// exit a loop
	sum = 0
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip odd numbers
		}
		sum += i
	}
	fmt.Println("sum:", sum) // 6 (2+4)
}

func f(n int) (int, error) {
	// idiomatic
	if n == 10 {
		return n, nil
	}
	return 0, fmt.Errorf("error %d operation", 010)
}
