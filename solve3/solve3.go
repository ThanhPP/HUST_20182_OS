package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Lock of Fork
type Fork struct {
	sync.Mutex
}
type Lock struct {
	locker sync.Mutex
	isLock bool
}

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
	leftLock, rightLock *Lock
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
	// randomPause(10)
	if p.leftLock.isLock && p.rightLock.isLock {
		p.leftLock.locker.Unlock()
		p.rightLock.locker.Unlock()
		p.rightLock.isLock = false
		p.rightLock.isLock = false
	} else {
		p.leftLock.locker.Lock()
		p.rightLock.locker.Lock()
		p.rightLock.isLock = true
		p.rightLock.isLock = true
		say("hungry", p.id)
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, p.id)
		p.leftFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, p.id)
		fmt.Printf("Philosopher #%d is taking Fork %d \n", p.id, ((p.id + 1) % 2))
		p.rightFork.Lock()
		fmt.Printf("Philosopher #%d is taked Fork %d  completely\n", p.id, ((p.id + 1) % 2))
		fmt.Printf("Philosopher #%d taked Fork completely \n", p.id)

		say("eating", p.id)
		// randomPause(10)
		say("Finished eating", p.id)
		fmt.Printf("Philosopher #%d is returning Fork \n", p.id)
		p.leftFork.Unlock()
		p.rightFork.Unlock()
		fmt.Printf("Philosopher #%d returned Fork completely \n", p.id)
		p.dine()
	}
}

var wg sync.WaitGroup

func main() {
	num := 2

	fork := make([]*Fork, num)
	lock := make([]*Lock, num)
	for i := 0; i < num; i++ {
		fork[i] = new(Fork)
		lock[i] = new(Lock)
	}

	philosopher := make([]*Philosopher, num)
	for i := 0; i < num; i++ {
		philosopher[i] = &Philosopher{
			id:        i,
			leftFork:  fork[i],
			rightFork: fork[(i+1)%num],
			leftLock:  lock[i],
			rightLock: lock[(i+1)%num],
		}
	}
	for i := 0; i < num; i++ {
		go philosopher[i].dine()
	}
	endless := make(chan int)
	<-endless
}
