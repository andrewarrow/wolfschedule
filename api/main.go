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

	for _, c := range listings {
		fmt.Printf("%s %s %s\n",
			//display.LeftAligned(c.Name, 10),
			display.LeftAligned(c.Symbol, 10),
			// CirculatingSupply
			// DateAdded
			// TotalSupply
			// MaxSupply
			// Symbol
			// DateAdded
			// NumMarketPairs
			// CMCRank
			display.LeftAligned(fmt.Sprintf("%0.2f", c.CirculatingSupply/1000000.0), 10),
			display.LeftAligned(fmt.Sprintf("%0.2f", c.TotalSupply/1000000.0), 10))
	}
}
