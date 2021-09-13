package main

import (
	"fmt"
	"os"
)

func calcNums(start, end int, f *os.File, done chan struct{}) {
	for i := start; i < end; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
	done <- struct{}{}
}

func main() {
	// write in file
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	noGoroutines := 10
	done := make(chan struct{})

	final := 16777215

	for i := 0; i < final; i = i + (final / noGoroutines) + 1 {
		paso := i + (final / noGoroutines)

		if paso > final {
			paso = final
		}
		go calcNums(i, paso, f, done)
	}

	// done
	doneNum := 0
	for doneNum < noGoroutines {
		<-done
		fmt.Println("done!")
		doneNum++
	}
	fmt.Println("Finished! :D")
}
