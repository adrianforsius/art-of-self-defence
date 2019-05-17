// +build ignore,OMIT

package main

import (
	"fmt"
)

// STOP OMIT
func main() {
	go func() { // HL
		fmt.Printf("Hello world, concurrently") // HL
	}() // HL
}
