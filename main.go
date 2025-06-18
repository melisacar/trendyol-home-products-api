package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
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
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ProductID         int    `json:"productGroupId"`
	CardType          string `json:"cardType"`
	CategoryHierarchy string `json:"categoryHierarchy"`
	CategoryId        int    `json:"categoryId"`
	CategoryName      string `json:"categoryName"`
	Url               string `json:"url"`
	MerchantId        int    `json:"merchantId"`
	CampaignName      string `json:"campaignName"`
	ItemNumber        int    `json:"itemNumber"`
	Brand             struct {
		BrandID   int    `json:"id"`
		BrandName string `json:"name"`
	}
	Price struct {
		SellingPrice    float64 `json:"sellingPrice"`
		OriginalPrice   float64 `json:"originalPrice"`
		DiscountedPrice float64 `json:"discountedPrice"`
		BuyingPrice     float64 `json:"buyingPrice"`
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
		fmt.Printf("Page %d: %d products have been added \n", page, len(pageResp.Result.Products))
	}

	// Result
	for _, product := range allProducts {
		fmt.Printf("procuct_id %d | name: %s | productGroupId: %d | cardType: %s | categoryHierarchy: %s | categoryId: %d | categoryName: %s | url: %s |  merchantId: %d | campaignName: %s | itemNumber: %d | brandId: %d | brandName: %s | sellingPrice: %.2f TL | originalPrice: %.2f | discountedPrice: %.2f | buyingPrice: %.2f | brandName: %s | categoryName: %s\n",
			product.ID, product.Name, product.ProductID, product.CardType, product.CategoryHierarchy, product.CategoryId, product.CategoryName, product.Url, product.MerchantId, product.CampaignName, product.ItemNumber, product.Brand.BrandID, product.Brand.BrandName, product.Price.SellingPrice, product.Price.OriginalPrice, product.Price.DiscountedPrice, product.Price.BuyingPrice, product.Brand.BrandName, product.CategoryName)
	}

	// Create excel file
	f := excelize.NewFile()
	sheet := "Products"
	f.NewSheet(sheet)

	today := time.Now().Format("2006-01-02") // YYYY-MM-DD

	// Headers
	headers := []string{
		"product_id", "name", "productGroupId", "cardType", "categoryHierarchy", "categoryId",
		"categoryName", "url", "merchantId", "campaignName", "itemNumber",
		"brand_id", "brand_name", "sellingPrice", "originalPrice", "discountedPrice", "buyingPrice",
		"marka", "yma_category", "scraped_date",
	}
	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, h)
	}

	// Add products
	for i, p := range allProducts {
		row := i + 2 // Row 1 is for headers
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.ProductID)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.CardType)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), p.CategoryHierarchy)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), p.CategoryId)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), p.CategoryName)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), p.Url)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), p.MerchantId)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), p.CampaignName)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), p.ItemNumber)
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), p.Brand.BrandID)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), p.Brand.BrandName)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), p.Price.SellingPrice)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), p.Price.OriginalPrice)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), p.Price.DiscountedPrice)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), p.Price.BuyingPrice)
		f.SetCellValue(sheet, fmt.Sprintf("R%d", row), p.Brand.BrandName)
		f.SetCellValue(sheet, fmt.Sprintf("S%d", row), p.CategoryName)
		f.SetCellValue(sheet, fmt.Sprintf("T%d", row), today)
	}

	// Set active sheet and save the file
	index, err := f.GetSheetIndex(sheet)
	if err != nil {
		log.Fatalf("Failed to get sheet index: %v", err)
	}
	f.SetActiveSheet(index)

	if err := f.SaveAs("trendyol_products_davlumbaz_formatted.xlsx"); err != nil {
		log.Fatalf("Failed to save Excel file: %v", err)
	}

	fmt.Println("Excel file created: trendyol_products_davlumbaz_formatted.xlsx")

}
