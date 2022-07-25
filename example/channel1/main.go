package main

import "fmt"

func main() {
	fmt.Println("hi channel")

	mych := make(chan string)

	go func() {
		mych <- "abc"
		mych <- "xyz"
		close(mych)
	}()

	for val := range mych {
		fmt.Println(val)
	}
}
