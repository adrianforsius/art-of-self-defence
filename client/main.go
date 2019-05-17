package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "http://localhost:8081"

type Apple struct {
	ID    int    `json:"id"`
	Fresh bool   `json:"fresh"`
	Color string `json:"color"`
}

func main() {
	c := Fruit{
		&Client{
			Client: http.Client{
				Timeout: 3 * time.Second,
			},
		},
	}
	apples, err := c.FreshApples()
	if err != nil {
		log.Printf("error getting apples %s", err)
	}
	fmt.Print("My resp %v", apples)
}

type Client struct {
	http.Client
}

func (c *Client) Apples() ([]Apple, error) {
	resp, err := http.Get(url + "/apples")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var a []Apple
	err = json.Unmarshal(body, &a)
	return a, err
}

func (c *Client) Apple(id int) (Apple, error) {
	var a Apple
	resp, err := http.Get(url + "/apple")
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}
	err = json.Unmarshal(body, &a)
	return a, err
}

type Fruit struct {
	API *Client
}

type Result struct {
	Apple Apple
	Error error
}

func (f Fruit) FreshApples() ([]Apple, error) {
	// START1 OMIT
	// List all the apples, remeber this is only a list of id's
	apples, err := f.API.Apples() // HL
	if err != nil {
		return nil, err
	}
	// STOP1 OMIT

	// START2 OMIT
	// Make apple channel
	appleChannel := make(chan Result)
	defer close(appleChannel)

	for _, apple := range apples {
		// Make a request for each apple in the list
		// We use apple pointer here to be able to signal failure over creating a custom type
		go func(id int, appleChannel chan<- Result) { // HL
			apple, err := f.API.Apple(id) // HL
			appleChannel <- Result{
				Apple: apple,
				Error: err,
			}
		}(apple.ID, appleChannel)
	}
	// STOP2 OMIT

	var freshApples []Apple
	// START3 OMIT
	count := 0
	// Loop over all go routines making sure we don't leave go routines behind
	for count < len(apples) {
		result := <-appleChannel // HL
		apple := result.Apple
		if result.Error != nil {
			err = result.Error
		}
		if apple.Fresh {
			// If the apple is fresh add it to our fresh apples
			freshApples = append(freshApples, apple) // HL
		}
		count++
	}
	// STOP3 OMIT
	if err != nil {
		return nil, err
	}

	return freshApples, err
}
