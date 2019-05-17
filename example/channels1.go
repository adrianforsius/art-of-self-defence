// +build ignore,OMIT

package main

func main() {
	var apple Apple
	// START1 OMIT
	// Apple definition
	type Apple struct { // HL
		ID    int  // HL
		Fresh bool // HL
	} // HL
	// STOP1 OMIT

	// START2 OMIT
	// Declaring and initializing.
	c := make(chan Apple) // HL
	// STOP2 OMIT

	// START3 OMIT
	// Send an fresh apple on the channel
	c <- Apple{ID: 1, Fresh: true} // HL
	// STOP3 OMIT

	// START4 OMIT
	// Receiving from a channel.
	// The "arrow" indicates the direction of data flow.
	recievedApple = <-c // HL
	// STOP4 OMIT

	// START5 OMIT
	// We can even send a channel on a channel
	channelChannel := make(chan chan Apple) // HL
	// STOP5 OMIT

	_ = recievedApple
}
