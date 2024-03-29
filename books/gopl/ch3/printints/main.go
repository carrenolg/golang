package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	// fmt.Println(buf)
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	// fmt.Println(buf)
	return buf.String()
}

func main() {
	// print ints
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
