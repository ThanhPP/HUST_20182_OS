package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//Lock of Fork
type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
func say(str string, id int) {
	fmt.Printf("Philosopher #%d is %s \n", id, str)
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func (p Philosopher) dine() {
	say("thinking", p.id)
	randomPause(10)
	say("hungry", p.id)
	if p.id%2 == 0 {
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, p.id)
		p.leftFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, p.id)
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, ((p.id + 1) % 2))
		p.rightFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, ((p.id + 1) % 2))
		fmt.Printf("Philosopher #%d taked Fork completely \n", p.id)

		say("eating", p.id)
		randomPause(10)
		say("Finished eating", p.id)
		fmt.Printf("Philosopher #%d is returning Fork \n", p.id)
		p.leftFork.Unlock()
		p.rightFork.Unlock()
		fmt.Printf("Philosopher #%d returned Fork completely \n", p.id)
	} else {
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, ((p.id + 1) % 2))
		p.rightFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, ((p.id + 1) % 2))
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, p.id)
		p.leftFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, p.id)
		fmt.Printf("Philosopher #%d taked Fork completely \n", p.id)

		say("eating", p.id)
		randomPause(10)
		say("Finished eating", p.id)
		fmt.Printf("Philosopher #%d is returning Fork \n", p.id)
		p.rightFork.Unlock()
		p.leftFork.Unlock()
		fmt.Printf("Philosopher #%d returned Fork completely \n", p.id)
	}
	p.dine()
}

var wg sync.WaitGroup

func main() {
	num := 2
	runtime.GOMAXPROCS(num)

	fork := make([]*Fork, num)
	for i := 0; i < num; i++ {
		fork[i] = new(Fork)
	}

	philosopher := make([]*Philosopher, num)
	for i := 0; i < num; i++ {
		philosopher[i] = &Philosopher{
			id:        i,
			leftFork:  fork[i],
			rightFork: fork[(i+1)%num],
		}
		go philosopher[i].dine()
	}
	for i := 0; i < num; i++ {

	}
	endless := make(chan int)
	<-endless
}
