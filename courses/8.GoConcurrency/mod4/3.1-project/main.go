package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopStick
	host            chan struct{}
}

func (p Philo) eat() {
	for range 3 {
		// request permission from host to eat
		p.host <- struct{}{}
		p.rightCS.Lock()
		p.leftCS.Lock()
		fmt.Printf("starting to eat [%d]\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		// release permission back to host
		<-p.host
		fmt.Printf("finishing eating [%d]\n", p.id)
	}
	wg.Done()
}

func main() {

	ChopSticks := make([]*ChopStick, 5)

	for i := range 5 {
		ChopSticks[i] = new(ChopStick)
	}

	// host (arbitrator) allows at most 2 philosophers to eat concurrently
	host := make(chan struct{}, 2)

	Philosophers := make([]*Philo, 5)
	for i := range 5 {
		Philosophers[i] = &Philo{
			id:      i + 1,
			leftCS:  ChopSticks[i],
			rightCS: ChopSticks[(i+1)%5],
			host:    host,
		}
	}

	wg.Add(5)

	for i := range 5 {
		go Philosophers[i].eat()
	}

	wg.Wait()
}
