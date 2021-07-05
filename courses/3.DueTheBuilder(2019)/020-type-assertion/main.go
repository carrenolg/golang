package main

import (
	"fmt"
	"time"
)

func main() {
	// type assertions
	v := Video{}
	a := Audio{}
	Info(a)
	Info(v)
	Print(15)
	Print(5.5)
	m := Message{"Hello"}
	Print(m)
}

type Audio struct {
	// fields
}

type Video struct {
	// fields
}

type Message struct {
	text string
}

func (m Message) String() string {
	return fmt.Sprintf("The message is: %v", m.text)
}

func (a Audio) Title() string           { return "Title" }
func (a Audio) Duration() time.Duration { return 3600 }
func (a Audio) Fortmat() string         { return "MP3" }
func (a Audio) Author() string          { return "Author" }

func (v Video) Title() string           { return "Title" }
func (v Video) Duration() time.Duration { return 10600 }
func (v Video) Fortmat() string         { return "MP4" }
func (v Video) Director() string        { return "Director" }

type Media interface {
	Title() string
	Duration() time.Duration
	Fortmat() string
}

func Info(m Media) {
	fmt.Println(m.Title(), m.Duration(), m.Fortmat())
	// type assertion v.(type)
	a, ok := m.(Audio)
	if ok {
		fmt.Println(a.Author())
	}
	v, ok := m.(Video)
	if ok {
		fmt.Println(v.Director())
	}
	// type switch
	switch t := m.(type) {
	case Audio:
		fmt.Println(t.Author())
	case Video:
		fmt.Println(t.Director())
	}
}

// empty interface
func Print(i interface{}) {
	switch t := i.(type) {
	case int8, int16, int32, int64, int:
		fmt.Printf("%d\n", t)
	case float32, float64:
		fmt.Printf("%f\n", t)
	case fmt.Stringer:
		//fmt.Println(t.String())
		fmt.Println(t)
	default:
		fmt.Printf("%v\n", t) // general placeholder
	}
}
