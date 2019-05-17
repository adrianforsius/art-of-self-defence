// +build ignore,OMIT
package main

import "fmt"

func main() {
	// START1 OMIT
	// Make
	queue := make(chan string)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
	// STOP1 OMIT
}
