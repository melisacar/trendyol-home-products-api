package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ApiResponse struct {
	IsSuccess  bool `json:"isSuccess"`
	StatusCode int  `json:"statusCode"`
	Error      any  `json:"error"`
	Result     struct {
		SlpName    string    `json:"slpName"`
		Products   []Product `json:"products"`
		TotalCount int       `json:"totalCount"`
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
	baseURL := "https://apigw.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/davlumbaz-x-c103627"
	pageSize := 24

	allProducts := []Product{}

	// First page request
	firstPage := 1
	url := fmt.Sprintf("%s?pi=%d&culture=tr-TR&userGenderId=1&channelId=1", baseURL, firstPage)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request error on first page: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response on first page: %v", err)
	}

	var apiResp ApiResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatalf("Error decoding JSON on first page: %v", err)
	}

	allProducts = append(allProducts, apiResp.Result.Products...)

	totalCount := apiResp.Result.TotalCount
	totalPages := (totalCount + pageSize - 1) / pageSize // ceil

	fmt.Printf("Total product number: %d\n", totalCount)
	fmt.Printf("Total page number: %d\n", totalPages)

	//
	for page := 2; page <= totalPages; page++ {
		url := fmt.Sprintf("%s?pi=%d&culture=tr-TR&userGenderId=1&channelId=1", baseURL, page)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Request error on page %d: %v", page, err)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error reading response on page %d: %v", page, err)
			continue
		}

		var pageResp ApiResponse
		err = json.Unmarshal(body, &pageResp)
		if err != nil {
			log.Printf("Error decoding JSON on page %d: %v", page, err)
			continue
		}

		allProducts = append(allProducts, pageResp.Result.Products...)
		fmt.Printf("Page %d: %d ürün eklendi\n", page, len(pageResp.Result.Products))
	}

	// Result
	for _, product := range allProducts {
		fmt.Printf("ID: %d | Name: %s | Price: %.2f TL\n", product.ID, product.Name, product.Price.SellingPrice)
	}
}
