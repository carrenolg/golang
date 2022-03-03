package main

import (
	"errors"
	"fmt"
	"log"
)

type MyBuffer interface {
	insert(element int)
	remove() (int, bool)
}

type buffer struct {
	stack []int
	btype string
}

// arrar
//type stack []buffer

func NewBuffer(btype string) (*buffer, error) {
	if btype != "FIFO" && btype != "LIFO" {
		return nil, errors.New("invalid btype parameter")
	}
	var stack []int
	return &buffer{btype: btype, stack: stack}, nil
}

func (b *buffer) insert(element int) {
	b.stack = append(b.stack, element)
}

func (b *buffer) remove() (int, bool) {
	var removed int
	var status bool
	switch b.btype {
	case "FIFO":
		removed = b.stack[0]
		b.stack = b.stack[1:]
		status = true
	case "LIFO":
		removed = b.stack[len(b.stack)-1]
		b.stack = b.stack[:len(b.stack)-1]
		status = true
	default:
		removed = -1
		status = false
	}
	return removed, status
}

func main() {
	// FIFO
	newBuffer, err := NewBuffer("FIFO")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	newBuffer.insert(10)
	newBuffer.insert(20)
	newBuffer.insert(30)
	newBuffer.insert(40)
	fmt.Println(newBuffer.stack)
	newBuffer.remove()
	newBuffer.remove()
	fmt.Println(newBuffer.stack)
	// LIFO
	secondBuffer, err := NewBuffer("LIFO")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	secondBuffer.insert(10)
	secondBuffer.insert(20)
	secondBuffer.insert(30)
	secondBuffer.insert(40)
	fmt.Println(secondBuffer.stack)
	secondBuffer.remove()
	secondBuffer.remove()
	fmt.Println(secondBuffer.stack)
}
