package main

import (
	"fmt"
	"sync"
)

func main() {
	fork := new(sync.Mutex)
	fork.Lock()
	fmt.Println("a")
	fork.Lock()
}
