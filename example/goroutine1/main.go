//iscuss how to create and work on multiple goroutines with the help of an example:

// Go program to illustrate Multiple Goroutines
package main

import (
	"fmt"
	"time"
)

func Aname1(ch chan string) {

	arr1 := [4]string{"a", "b"}

	for t1 := 0; t1 <= 1; t1++ {
		ch <- arr1[t1]
		//time.Sleep(150 * time.Millisecond)
		//fmt.Printf("%s\n", arr1[t1])
	}
	close(ch)
}

// For goroutine 1
func Aname(ch chan string) {

	arr1 := [4]string{"Rohit", "Suman"}

	for t1 := 0; t1 <= 1; t1++ {
		ch <- arr1[t1]
		//time.Sleep(150 * time.Millisecond)
		//fmt.Printf("%s\n", arr1[t1])
	}
	close(ch)
}
func display(ch chan string) {
	fmt.Println("jjjjjj")
	for val := range ch {
		fmt.Println(val)
	}
}

// For goroutine 2
func Aid() {

	arr2 := [4]int{300, 301}

	for t2 := 0; t2 <= 1; t2++ {

		//time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

// Main function
func main() {
	mych := make(chan string)
	fmt.Println("!...Main Go-routine Start...!")
	go Aname(mych)
	//go Aid()
	//go Aname1(mych)
	//go display(mych)
	time.Sleep(500 * time.Millisecond)
	//close(mych)
	//fmt.Println(<-mych)
	for val := range mych {
		fmt.Println(val)
	}
	for val := range mych {
		fmt.Println(val)
	}
	fmt.Println("\n!...Main Go-routine End...!")
}
