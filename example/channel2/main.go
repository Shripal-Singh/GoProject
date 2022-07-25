package main

import "fmt"

func mychannel(ch chan int) {
	fmt.Println(254 + <-ch)

}
func main() {

	mych := make(chan int)

	go mychannel(mych)
	mych <- 23

	//fmt.Println(<-mych)
}
