package main

import "fmt"

// function 1
func portal1(channel1 chan string) {
	for i := 0; i <= 3; i++ {
		channel1 <- "Welcome to channel 1"
	}

}

// function 2
func portal2(channel2 chan string) {
	channel2 <- "Welcome to channel 2"
}

// main function
func main() {

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	// the choice of selection
	// of case is random
	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
}
