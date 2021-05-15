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

	template := `<tr>
<td>%s</td><td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td> <td>0.0</td> <td>0.0</td> <td>0.0</td> <td>0.0</td>
</tr>`

	client := cmc.NewClient(&cmc.Config{ProAPIKey: api})

	listings, _ := client.Cryptocurrency.
		LatestListings(&cmc.ListingOptions{
			Limit: 1000,
		})

	for _, c := range listings {
		if c.Symbol != "ADA" && c.Symbol != "ALGO" &&
			c.Symbol != "IOTA" && c.Symbol != "NANO" &&
			c.Symbol != "EGLD" && c.Symbol != "CELO" &&
			c.Symbol != "ATOM" && c.Symbol != "LUNA" &&
			c.Symbol != "BTC" && c.Symbol != "ETH" &&
			c.Symbol != "DOGE" && c.Symbol != "XLM" &&
			c.Symbol != "VET" &&
			c.Symbol != "QTUM" &&
			c.Symbol != "ONE" &&
			c.Symbol != "KAVA" && c.Symbol != "BNB" && c.Symbol != "KSM" &&
			c.Symbol != "XTZ" && c.Symbol != "DOT" {
			continue
		}
		usd := c.Quote["USD"].Price
		mcap := (c.CirculatingSupply / 1000000000.0) * usd

		html := fmt.Sprintf(template, c.Name, display.LeftAligned(c.DateAdded, 4),
			c.Symbol,
			fmt.Sprintf("%0.2f", mcap),
			display.LeftAligned(c.NumMarketPairs, 10))
		fmt.Println(html)
	}
}
func main2() {
	api := os.Getenv("CMC")

	client := cmc.NewClient(&cmc.Config{ProAPIKey: api})

	listings, _ := client.Cryptocurrency.
		LatestListings(&cmc.ListingOptions{
			Limit: 1000,
		})

	fmt.Printf("%s %s %s %s%s%s %s %s\n",
		display.LeftAligned("", 10),
		display.LeftAligned("", 8),
		display.LeftAligned("Pairs", 10),
		display.LeftAligned("Max", 10),
		display.LeftAligned("Cir", 12),
		display.LeftAligned("Total", 12),
		display.LeftAligned("USD", 10),
		display.LeftAligned("Cap", 10))
	for _, c := range listings {
		if c.Symbol == "SAFEMOON" || c.Symbol == "ELON" {
			continue
		}
		usd := c.Quote["USD"].Price
		mcap := (c.CirculatingSupply / 1000000000.0) * usd
		fmt.Printf("%s %s %s %s%s%s %s %s\n",
			//display.LeftAligned(c.Name, 10),
			display.LeftAligned(c.DateAdded, 10),
			display.LeftAligned(c.Symbol, 8),
			display.LeftAligned(c.NumMarketPairs, 10),
			// CMCRank
			display.LeftAligned(fmt.Sprintf("%0.2f", c.MaxSupply/1000000000.0), 10),
			display.LeftAligned(fmt.Sprintf("%0.2f", c.CirculatingSupply/1000000000.0), 12),
			display.LeftAligned(fmt.Sprintf("%0.2f", c.TotalSupply/1000000000.0), 12),
			display.LeftAligned(fmt.Sprintf("%0.2f", usd), 10),
			display.LeftAligned(fmt.Sprintf("%0.2f", mcap), 10))
	}
}
