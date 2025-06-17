package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://apigw.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/davlumbaz-x-c103627?pi=1&culture=tr-TR&userGenderId=1&channelId=1"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error on reading the response", err)
	}

	fmt.Println(string(body))
}
