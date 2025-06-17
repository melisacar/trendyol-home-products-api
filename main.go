package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ApiResponse struct {
	Result struct {
		Products []Product `json:"products"`
	} `json:"result"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price struct {
		SellingPrice float64 `json:"sellingPrice"`
	} `json:"price"`
}

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

	var apiResp ApiResponse

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	for _, product := range apiResp.Result.Products {
		fmt.Printf("Product ID: %d - Name: %s - Price: %.2f\n", product.ID, product.Name, product.Price.SellingPrice)
	}
}
