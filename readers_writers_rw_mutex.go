package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUMBER_OF_READERS = 5
const NUMBER_OF_WRITERS = 5
const MAX_DURATION = 5

var (
	mutex sync.RWMutex
	data  int32
	group sync.WaitGroup
)

func read_unit(i int) {
	n := rand.Intn(MAX_DURATION)
	fmt.Println("Reader #", i, " read: ", data, " and sleep for: ", n, " seconds")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Reader #", i, " finished")
}

func write_unit(i int) {
	data = rand.Int31n(10)
	n := rand.Intn(MAX_DURATION)
	fmt.Println("Writer #", i, " wrote: ", data, " and sleep for: ", n, " seconds")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Writer #", i, " finished")
}

func readers() {
	for i := 1; i <= NUMBER_OF_READERS; i++ {
		// reader goroutine
		go func(i int) {
			group.Add(1)

			mutex.RLock()
			read_unit(i)
			mutex.RUnlock()

			group.Done()
		}(i)
	}
}

func writers() {
	for i := 1; i <= NUMBER_OF_WRITERS; i++ {
		// writer goroutine
		go func(i int) {
			group.Add(1)

			mutex.Lock()
			write_unit(i)
			mutex.Unlock()

			group.Done()
		}(i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data = 1

	writers()
	readers()

	group.Wait()
	fmt.Println("Data: ", data)
}
