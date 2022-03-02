package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants
const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

// Queue
type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New
func (q *Queue) New() {
	q.message = make(chan int)
	q.queuePass = make(chan int)
	q.queueTicket = make(chan int)

	// selecting the message based on its type
	go func() {
		var message int
		for {
			select {
			case message = <-q.message:
				switch message {
				case messagePassStart:
					q.waitPass++
				case messagePassEnd:
					q.playPass = false
				case messageTicketStart:
					q.waitTicket++
				case messageTicketEnd:
					q.playTicket = false
				}
				if q.waitPass > 0 && q.waitTicket > 0 && !q.playPass && !q.playTicket {
					q.playPass = true
					q.playTicket = true
					q.waitTicket--
					q.waitPass--
					q.queuePass <- 1
					q.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (queue *Queue) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

// EndTicketIssue ends the ticket issue
func (queue *Queue) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

//StartPass ends the Pass queue
func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

//EndPass ends the Pass queue
func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

//ticketIssue starts and ends the ticket issue
func ticketIssue(queue *Queue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		queue.EndTicketIssue()
	}
}

//passenger method starts and ends the pass queue
func passenger(queue *Queue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartPass()
		fmt.Println("  Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")

		queue.EndPass()
	}
}

// main method
func main() {
	var queue *Queue = &Queue{}
	queue.New()
	fmt.Println(queue)
	var i int
	for i = 0; i < 10; i++ {
		go passenger(queue)
	}
	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(queue)
	}
	select {}
}
