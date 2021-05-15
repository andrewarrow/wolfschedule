package main

import (
	"fmt"

	"api/display"
	"os"

	//cmc "github.com/coincircle/go-coinmarketcap"
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	api := os.Getenv("CMC")

	client := cmc.NewClient(&cmc.Config{ProAPIKey: api})

	listings, _ := client.Cryptocurrency.
		LatestListings(&cmc.ListingOptions{
			Limit: 10,
		})

	fmt.Printf("%s %s %s %s %s %s %s\n",
		display.LeftAligned("", 10),
		display.LeftAligned("", 10),
		display.LeftAligned("Pairs", 10),
		display.LeftAligned("Max", 10),
		display.LeftAligned("Circulating", 15),
		display.LeftAligned("Total", 10),
		display.LeftAligned("USD", 10))
	for _, c := range listings {
		usd := c.Quote["USD"].Price
		fmt.Printf("%s %s %s %s %s %s %s\n",
			//display.LeftAligned(c.Name, 10),
			display.LeftAligned(c.DateAdded, 10),
			display.LeftAligned(c.Symbol, 10),
			display.LeftAligned(c.NumMarketPairs, 10),
			// CMCRank
			display.LeftAligned(fmt.Sprintf("%0.2f", c.MaxSupply/1000000.0), 10),
			display.LeftAligned(fmt.Sprintf("%0.2f", c.CirculatingSupply/1000000.0), 15),
			display.LeftAligned(fmt.Sprintf("%0.2f", c.TotalSupply/1000000.0), 10),
			display.LeftAligned(fmt.Sprintf("%0.2f", usd), 10))
	}
}
