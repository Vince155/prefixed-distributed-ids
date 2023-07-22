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

	flag := true

	for flag {
		counter++
	}

	fmt.Println("Hello, world")
}
