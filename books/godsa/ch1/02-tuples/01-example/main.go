package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	var squar int
	var cube int
	squar, cube = powerSeries(3)
	fmt.Println("Squar", squar, "Cube", cube)

}
