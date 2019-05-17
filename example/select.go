// +build ignore,OMIT

package main

import "fmt"

func main() {
	var c1, c2, c3 chan int
	// START0 OMIT
	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2: // HL
		fmt.Printf("received %v from c2\n", v1)
	case c3 <- 23: // HL
		fmt.Printf("sent %v to c3\n", 23)
	}
	// STOP0 OMIT
}
