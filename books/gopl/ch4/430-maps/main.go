package main

import "fmt"

func main() {
	// maps is a reference a hash table
	m := make(map[float64]string)
	m[2.1] = "dos.uno"
	m[2.2] = "dos.dos"
	fmt.Println(m)
}
