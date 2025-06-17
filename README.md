# Trendyol Product Scraper - Davlumbaz Category

This Go program scrapes product data from the **Davlumbaz** category on [Trendyol](https://www.trendyol.com), retrieves all paginated product listings, and exports the data into a well-structured Excel file.

---

## Features

- Sends paginated GET requests to Trendyol's public API endpoint.
- Parses and aggregates all products in the selected category.
- Extracts key fields: `ID`, `Name`, `Selling Price`.
- Saves the data to a structured `.xlsx` Excel file using the [`excelize`](https://github.com/xuri/excelize) library.

---

## Output

The final Excel file: `trendyol_products_davlumbaz.xlsx`  
Includes the following columns:

| ID  | Name           | Price (TL) |
|-----|----------------|------------|
| 123 | Example Name 1 | 599.90     |
| 456 | Example Name 2 | 749.00     |

---

## Getting Started

### 1. Clone this repository (if needed)

```bash
git clone https://github.com/your-username/trendyol-product-scraper.git
cd trendyol-product-scraper
```

### 2. Install dependencies

This project uses [Go modules](https://blog.golang.org/using-go-modules).

```bash
go mod init trendyol-home-products-api
go get github.com/xuri/excelize/v2
```

### 3. Run the script

```bash
go run main.go
```

---

## How It Works

- The script sends a request to the Trendyol API endpoint for the *Davlumbaz* category.
- It determines the total number of products and calculates how many pages need to be fetched.
- Then it loops through each page and aggregates the product data.
- After collecting all products, it writes them into a formatted Excel file.

---

## Dependencies

- `net/http` – for sending GET requests  
- `encoding/json` – for parsing the JSON responses  
- `github.com/xuri/excelize/v2` – for Excel file generation

---

## API Endpoint Used

```txt
https://apigw.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/davlumbaz-x-c103627
```

Query parameters such as `pi` (page index), `culture`, `userGenderId`, and `channelId` are used to paginate through product listings.

---

## Sample Output Log

```txt
Total product number: 1080  
Total page number: 45  
Page 2: 24 have been added
...  
Excel file created: trendyol_products_davlumbaz.xlsx
```

---

## License

This project is licensed under the **MIT License**.  

---