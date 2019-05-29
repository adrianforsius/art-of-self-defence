// +build ignore,OMIT
package main

import "fmt"

func main() {
	// START1 OMIT
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
	// STOP1 OMIT
}
