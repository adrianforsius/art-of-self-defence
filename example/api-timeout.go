// +build ignore,OMIT
package main

import "time"

func main() {
	// START1 OMIT
	for {
		select {
		case <-orig.Context.Done():
			request.Context.Done()
		// Generator (return recieve channel)
		case response := <-API.Fruit():
			// Continue
		}
	}
	// STOP1 OMIT

	// START2 OMIT
	for {
		select {
		case <-FanIn:
			// Continue
		case time.After(2 * time.Second):
			req.Context.Done()
		}
	}
	// STOP2 OMIT
}
