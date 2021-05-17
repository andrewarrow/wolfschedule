package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"text/template"

	"os"
	//cmc "github.com/coincircle/go-coinmarketcap"
	//cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

type Replacer struct {
	Caps    map[string]string
	Vol24   map[string]string
	Change1 map[string]string
}

func MakeHtml(r Replacer) {
	b2, _ := ioutil.ReadFile("stake.list.tmpl")
	blob := string(b2)
	t := template.Must(template.New("tmpl").
		Parse(blob))
	var buff bytes.Buffer
	t.Execute(&buff, r)
	ioutil.WriteFile("../marketing/stake.list.html", buff.Bytes(), 0755)
}

type CMC struct {
	Name           string
	Symbol         string
	NumMarketPairs string `json:"num_market_pairs"`
	DateAdded      string `json:"date_added"`
	Tags           []string
	MaxSupply      string `json:"max_supply"`
	Circulating    string `json:"circulating_supply"`
	TotalSupply    string `json:"total_supply"`
	Platform       string
	Quote          map[string]interface{}

	/*
	   "quote": {
	     "USD": {
	       "price": 43187.585412696455,
	       "volume_24h": 82493265835.38838,
	       "percent_change_1h": -1.24368813,
	       "percent_change_24h": -8.23503955,
	       "percent_change_7d": -25.30824614,
	       "percent_change_30d": -29.37895374,
	       "percent_change_60d": -27.89324191,
	       "percent_change_90d": -11.50405066,
	       "market_cap": 808143632402.0536,
	*/
}

type CMCHolder struct {
	Data []CMC `json:"data"`
}
type USD struct {
}

func main() {
	pat := os.Getenv("CMC")
	jsonString := DoGet(pat, "v1/cryptocurrency/listings/latest")
	var cmcHolder CMCHolder
	json.Unmarshal([]byte(jsonString), &cmcHolder)
	fmt.Println(cmcHolder)

	//template := `<tr>
	//<td>%s</td><td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td> <td>0.0</td> <td>0.0</td> <td>0.0</td> <td>0.0</td>
	//</tr>`
	/*
			client := cmc.NewClient(&cmc.Config{ProAPIKey: api})

			listings, _ := client.Cryptocurrency.
				LatestListings(&cmc.ListingOptions{
					Limit: 1000,
				})
		r := Replacer{}
		r.Caps = map[string]string{}
		r.Vol24 = map[string]string{}
		r.Change1 = map[string]string{}
		for _, c := range listings {
			if c.Symbol != "ADA" && c.Symbol != "ALGO" &&
				c.Symbol != "MIOTA" && c.Symbol != "NANO" &&
				c.Symbol != "EGLD" && c.Symbol != "CELO" &&
				c.Symbol != "ATOM" && c.Symbol != "LUNA" &&
				c.Symbol != "BTC" && c.Symbol != "ETH" &&
				c.Symbol != "DOGE" && c.Symbol != "XLM" &&
				c.Symbol != "VET" &&
				c.Symbol != "AVAX" &&
				c.Symbol != "MATIC" &&
				c.Symbol != "QTUM" &&
				c.Symbol != "ONE" &&
				c.Symbol != "KAVA" && c.Symbol != "BNB" && c.Symbol != "KSM" &&
				c.Symbol != "XTZ" && c.Symbol != "DOT" {
				continue
			}
			usd := c.Quote["USD"].Price
			mcap := (c.CirculatingSupply / 1000000000.0) * usd

			r.Caps[c.Symbol] = fmt.Sprintf("%0.2f", mcap)
			//r.Vol24[c.Symbol]
			//r.Change1[c.Symbol]

			fmt.Printf("%+v\n", c)

				html := fmt.Sprintf(template, c.Name, display.LeftAligned(c.DateAdded, 4),
					c.Symbol,
					fmt.Sprintf("%0.2f", mcap),
					display.LeftAligned(c.NumMarketPairs, 10))
				fmt.Println(html)
		}
	*/
	//MakeHtml(r)
}

/*
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
*/
