package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, x := range nums {
			out <- x
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()
	return out
}

func main() {
	for x := range sq(gen(2, 3, 4, 5)) {
		fmt.Println(x)
	}
}