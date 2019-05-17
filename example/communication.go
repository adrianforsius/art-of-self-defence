// +build ignore,OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	c := make(chan string)
	helloWorld := func(c chan string) {
		c <- "hello world" // HL
	}
	go helloWorld(c)
	message := <-c // HL
	fmt.Printf("Message: %s", message)
	// STOP1 OMIT
}
