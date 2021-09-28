package main

import (
	"fmt"
	"os"
	"strings"
)

func calcNums(start, end int, out chan string) {
	var buf strings.Builder
	for i := start; i < end; i++ {
		fmt.Fprintf(&buf, "%06x\n", i)
	}
	out <- buf.String() // send
	// doneCalc <- struct{}{}
}

func writeFile(in chan string, f *os.File, doneWrite chan struct{}) {
	// receive
	for s := range in {
		_, err := f.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
	doneWrite <- struct{}{}
}

func main() {
	// write in file
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numbers := make(chan string)
	noGoroutines := 10
	// doneCalc := make(chan struct{})
	doneWrite := make(chan struct{})

	final := 16777215

	for i := 0; i < final; i = i + (final / noGoroutines) + 1 {
		paso := i + (final / noGoroutines)

		if paso > final {
			paso = final
		}
		go calcNums(i, paso, numbers)
		go writeFile(numbers, f, doneWrite)
	}
	/*for countdown := noGoroutines; countdown > 0; countdown-- {
		<-doneCalc // in each loop a goroutine finished!
		fmt.Println("Finished Calc!")
	}
	close(numbers)
	<-doneWrite*/
	<-doneWrite
	fmt.Println("Program exit :D")
}
