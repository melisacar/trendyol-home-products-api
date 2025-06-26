# Trendyol Product Scraper (Price Range Based)

This Go program scrapes product data from the **Tencere** (sample) category on [Trendyol](https://www.trendyol.com), based on a **price range filter**, retrieves paginated product listings, and exports the data into a structured Excel file.

---

## Features

- Sends GET requests to Trendyol's public API with price filter parameters.
- Paginates through all products in the selected category and range.
- Extracts key product fields including price, brand, category, and more.
- Exports all collected data to a well-formatted `.xlsx` Excel file using [`excelize`](https://github.com/xuri/excelize).

---

## API Endpoint Used

```bash
https://apigw.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/tencere-x-c1191?culture=tr-TR&userGenderId=1&channelId=1&prc=15000-*
```

- `pi`: Page index  
- `prc`: Price range (`e.g., 0-5000`, `10000-*`, etc.)  
- Other parameters: `culture`, `userGenderId`, and `channelId` are static.

---

## How It Works

- Sends a request to the Trendyol API using the selected category and price range.
- Retrieves total product and page count.
- Iterates through all available pages, collecting product data.
- Exports the entire dataset into an Excel spreadsheet.

---

## Output

The generated file:

```bash
trendyol_products_tencere_seti_15000_end.xlsx
```

The Excel file contains columns such as:

| product_id | name | productGroupId | cardType | categoryHierarchy | categoryId | categoryName | url | merchantId | campaignName | itemNumber | brand_id | brand_name | sellingPrice | originalPrice | discountedPrice | buyingPrice | scraped_date |
|------------|------|----------------|----------|--------------------|-------------|---------------|-----|-------------|---------------|-------------|-----------|-------------|----------------|----------------|------------------|--------------|----------------|

---

## Getting Started

### 1. Clone this repository (optional)

```bash
git clone https://github.com/melisacar/trendyol-home-products-api.git
cd trendyol-home-products-api
```

### 2. Install dependencies

```bash
go mod init trendyol-tencere-scraper
go get github.com/xuri/excelize/v2
```

### 3. Run the script

```bash
go run main.go
```

You can also redirect output logs to a file:

```bash
go run main.go > output.txt
```

---

## Sample Output Log

```txt
Total product number: 864  
Total page number: 36  
Page 1: 24 products have been added
Page 2: 24 products have been added
...
Excel file created: trendyol_products_tencere_seti_15000_end.xlsx
```

---

## Dependencies

- `net/http` – for sending HTTP requests  
- `encoding/json` – for decoding JSON responses  
- `github.com/xuri/excelize/v2` – for Excel file generation

---

## License

This project is licensed under the **MIT License**.

---