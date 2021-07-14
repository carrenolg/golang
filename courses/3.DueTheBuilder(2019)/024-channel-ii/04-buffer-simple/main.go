package main

import (
	"fmt"
	"time"
)

func toUnicode(character int) int {
	uni, err := fmt.Printf("%U\n", character)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}
	return uni
}

func main() {
	// buffer simple
	ch := make(chan int, 3)
	data := []int{21, 31, 24}

	// send to buffer
	go func() {
		for _, v := range data {
			ch <- toUnicode(v)
			fmt.Println("Saving data on channel buffer ...")
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()

	// get from buffer
	for d := range ch {
		fmt.Printf("Byte len %v\n", d)
		fmt.Println("Getting data from channel buffer ...")
		time.Sleep(500 * time.Millisecond)
	}

}
