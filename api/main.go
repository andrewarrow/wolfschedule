package main

import (
	"fmt"
	"os"

	//cmc "github.com/coincircle/go-coinmarketcap"
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	api := os.Getenv("CMC")
	fmt.Println("key", api)

	client := cmc.NewClient(&cmc.Config{ProAPIKey: api})

	listings, _ := client.Cryptocurrency.
		LatestListings(&cmc.ListingOptions{
			Limit: 10,
		})

	for _, listing := range listings {
		fmt.Println(listing.Name)
		fmt.Println(listing.Quote["USD"].Price)
	}
}
