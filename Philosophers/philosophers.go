package main

import (
	"fmt"
	"sync"
)

// Philosopher represents a philosopher in the dining philosophers problem.
type Philosopher struct {
	id      int
	meals   chan bool
	left    *sync.Mutex
	right   *sync.Mutex
}

// Dine implements the dining behavior of a philosopher.
func (p *Philosopher) Dine(philosophers []*Philosopher) {
	for {
		// Pick up chopsticks
		p.left.Lock()
		p.right.Lock()

		// Eat
		p.meals <- true

		// Put down chopsticks
		p.right.Unlock()
		p.left.Unlock()
	}
}

func main() {
	numPhilosophers := 5

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:    i,
			meals: make(chan bool, 3),
			left:  new(sync.Mutex),
			right: new(sync.Mutex),
		}
	}

	for i := 0; i < numPhilosophers; i++ {
		go func(p *Philosopher) {
			p.Dine(philosophers)
		}(philosophers[i])
	}

	for {
		for i, p := range philosophers {
			fmt.Printf("Philosopher %d has eaten %d meals\n", p.id, len(p.meals))
			if i == numPhilosophers-1 {
				fmt.Println()
			}
		}
	}
}
