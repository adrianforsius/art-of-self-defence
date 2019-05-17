package main

import (
	"fmt"
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
	greenApples, err := f.API.Apples("green")
	if err != nil {
		return nil, err
	}

	appleChannel := make(chan Apple, len(greenApples))
	defer close(appleChannel)

	for _, greenApple := range greenApples {
		go func(id string, appleChannel chan<- *PrefectApple) {
			apple, err := f.API.Apple(id)
			if err != nil {
				fmt.Printf("failed to inspect green apple: %v", err)
			}
			appleChannel <- apple
		}(greenApple.ID, appleChannel)
	}
	close(appleChannel)

	var prefectApples []Apple
	count := 0

	for apple := range appelChannel {
		if apple.Perfect {
			perfectApples = append(perfectApples, apple)
		}
	}

	return perfectApples, err
}
