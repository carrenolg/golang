package main

import (
	"fmt"
	str "strings" // Package Alias

	"example.com/carrenolg/packer/numbers"
	"example.com/carrenolg/packer/strings"
	"example.com/carrenolg/packer/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
