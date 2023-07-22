package main

import (
	"fmt"
	"prefixed-distributed-ids/number_generator"
)

func main() {
	counter, err := number_generator.InitializeCounter()

	if err != nil {
		panic(err)
	}

	for counter < 1000 {
		counter++
	}

	fmt.Println("Hello, world")
}
