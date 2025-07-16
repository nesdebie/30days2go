package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"

	// Importing the Colly package (here and in the Makefile) for web scraping
	"github.com/gocolly/colly"
)

type Product struct {
    Url, Image, Name, Price string
}


func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: ", os.Args[0])
		return
	}

	const allowedDomain = "www.scrapingcourse.com"

	collectorObject := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
	)

	var products []Product

	// Callback for when a product is found on this specific site
	// Is called for each product found on the page
	collectorObject.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{}
		
		product.Url = e.ChildAttr("a", "href")
        product.Image = e.ChildAttr("img", "src")
        product.Name = e.ChildText(".product-name")
        product.Price = e.ChildText(".price")

        products = append(products, product)
	})

	// Callback for when the scraping is done
	// Is called once all the pages have been scraped
    collectorObject.OnScraped(func(r *colly.Response) {
        file, err := os.Create(os.Args[0] + ".csv")
        if err != nil {
            log.Fatalln("Failed to create output CSV file", err)
        }
        defer file.Close()

        writer := csv.NewWriter(file)

        headers := []string{
            "Url",
            "Image",
            "Name",
            "Price",
        }
        writer.Write(headers)

        for _, product := range products {
            record := []string{
                product.Url,
                product.Image,
                product.Name,
                product.Price,
            }

            writer.Write(record)
        }
        defer writer.Flush()
    })

	// Visit the ecommerce page & start the scraping process
	err := collectorObject.Visit("https://" + allowedDomain + "/ecommerce")
	if err != nil {
		fmt.Println("Error visiting:", err)
	}
}
