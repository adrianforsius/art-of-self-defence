// +build ignore,OMIT

package main

import (
	"fmt"
	"time"
)

// STOP OMIT
func main() {
	// Alternatively define the function first
	helloWorld := func() {
		fmt.Printf("hello world, concurrenctly")
	}
	go helloWorld() // HL

	time.Sleep(time.Second * 2)
}
