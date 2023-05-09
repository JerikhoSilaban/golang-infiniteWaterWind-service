package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	// randomWaterWind()

	ticker := time.NewTicker(2 * time.Second)

	// Create a channel for tickers
	tickerChan := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChan:
				return

			case <-ticker.C:
				randomWaterWind()
			}
		}
	}()

	for {
		// Calling Sleep() method
		time.Sleep(10 * time.Second)

		// Calling Stop() method
		ticker.Stop()

		// Setting the value of channel
		tickerChan <- true

		// Printed when the ticker is turned off
		fmt.Println("Ticker is turned off!")
	}
}

func randomWaterWind() {
	// Prepare data to be sent
	data := map[string]interface{}{
		"water": h8HelperRand.RandomInt(1, 100),
		"wind":  h8HelperRand.RandomInt(1, 100),
	}

	// Convert data to JSON form
	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	// Prepare a post request with the desired url
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	// Send post requests
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	// Convert the response to a slice of bytes
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	}

	var response Response

	_ = json.Unmarshal([]byte(body), &response)
	exp, _ := json.MarshalIndent(Response{Water: response.Water, Wind: response.Wind}, "", " ")

	fmt.Println(string(exp))

	var status_water string
	var status_wind string

	if response.Water < 5 {
		status_water = "aman"
	} else if response.Water >= 6 && response.Water <= 8 {
		status_water = "siaga"
	} else {
		status_water = "bahaya"
	}

	if response.Wind < 6 {
		status_wind = "aman"
	} else if response.Wind >= 7 && response.Wind <= 15 {
		status_wind = "siaga"
	} else {
		status_wind = "bahaya"
	}

	fmt.Printf("Status water adalah %s dengan ketinggian %d meter\n", status_water, response.Water)
	fmt.Printf("Status wind adalah %s dengan kecepatan %d meter/s\n", status_wind, response.Wind)
}
