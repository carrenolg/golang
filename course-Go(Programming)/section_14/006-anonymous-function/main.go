package main

import "fmt"

func main() {
	// called
	hello("Gio")

	// anonymous function
	func() {
		fmt.Println("Hello from the anony func")
	}()

	// playing
	f := func() {
		fmt.Println("this is a anony func")
	}

	f()
	fmt.Printf("%T\n", f) // type: [func()]

}

func hello(name string) {
	fmt.Println("Hello, I'm", name)
}
