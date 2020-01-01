package main

import (
	"fmt"
)

var n int = 10

func main() {
	fmt.Println("Hello")
	c := make(chan int)

	numOfGenerators := 10

	for i := 0; i < numOfGenerators; i++ {
		go generate(c)
	}

	for j := 0; j < numOfGenerators*n; j++ {
		v, ok := <-c

		if !ok {
			fmt.Println("Not ok")
		}

		fmt.Printf("Iteration %d: value: %d\n", j, v)
	}

	fmt.Println("Goodbye")
}

func generate(c chan<- int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}
