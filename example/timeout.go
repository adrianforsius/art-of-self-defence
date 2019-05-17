// +build ignore,OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START1 OMIT
	c := make(chan string)
	go func(c chan string) {
		for {
			select {
			case s := <-c: // HL
				time.Sleep(time.Second)
				fmt.Println(s)
				c <- "ping" // HL
			}
		}
	}(c)
	go func(c chan string) {
		for {
			select {
			case s := <-c: // HL
				time.Sleep(time.Second)
				fmt.Println(s)
				c <- "pong" // HL
			}
		}
	}(c)
	// infinite loop
	c <- "ping" // HL
	// STOP1 OMIT

	// START2 OMIT
	select {
	case <-time.After(2 * time.Second):
		<-c
		fmt.Printf("exiting both")
		time.Sleep(time.Second)
	}
	// STOP2 OMIT
}
