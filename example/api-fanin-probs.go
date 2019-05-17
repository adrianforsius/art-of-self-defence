// +build ignore,OMIT
package main

import (
	"net/http"
)

type Client struct {
	http.Client
}

func (c *Client) Apples(color string) ([]Apple, error) {
	req := http.Request("localhost:8081")
}

type FruitService struct {
	API *Client
}

func (f FruitService) NiceApples(email string) ([]Apple, error) {
	// START1 OMIT
	// List all the apples, remeber this is only a list of id's
	apples, err := f.API.Apples()
	if err != nil {
		return nil, err
	}

	// STOP1 OMIT

	// START2 OMIT
	appleChannel := make(chan Apple)
	// STOP2 OMIT

	// START3 OMIT
	for _, apple := range apples {
		go func(id string, appleChannel chan<- Apple) {
			apple, _ := f.API.Apple(id) // HL
			// ...
			appleChannel <- apple
		}(apple.ID, appleChannel)
	}
	// STOP3 OMIT

	// START4 OMIT
	var freshApples []Apple
	count := 0
	for apple := range appelChannel {
		if apple.Fresh {
			freshApples = append(freshApples, apple)
		}
		if len(apples) == count { // HL
			close(appelChannel)
		}
		count++
	}
	// STOP4 OMIT

	return freshApples, err
}
