package main

import (
	"fmt"
	"time"
)

func main() {
	// type assertions
	v := Video{}
	Info(v)
	a := Audio{}
	Info(a)
}

type Audio struct {
	// fields
}

type Video struct {
	// fields
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
}
