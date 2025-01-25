package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main1() {
	client := &http.Client{Timeout: 10 * time.Second}
	for {
		resp, err := client.Get("10.128.0.2")
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		var body []byte
		resp.Body.Read(body)
		fmt.Println("Response from server: ", string(body))
		time.Sleep(5 * time.Second)
	}
}
