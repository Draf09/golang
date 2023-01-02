package main

import (
	"testing"
	"time"
	"sync"
)

func TestDine(t *testing.T) {
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
		allHaveEaten := true
		for _, p := range philosophers {
			if len(p.meals) == 0 {
				allHaveEaten = false
				break
			}
		}
		if allHaveEaten {
			return
		}
		
		time.Sleep(time.Second)
	}
}
