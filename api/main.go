package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"os"
	//cmc "github.com/coincircle/go-coinmarketcap"
	//cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

type Replacer struct {
	Caps    map[string]string
	Vol24   map[string]string
	Change1 map[string]string
	Things  []CMC
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
	Slug           string `json:"slug"`
	Symbol         string
	NumMarketPairs string `json:"num_market_pairs"`
	DateAdded      string `json:"date_added"`
	Tags           []string
	MaxSupply      float64 `json:"max_supply"`
	Circulating    float64 `json:"circulating_supply"`
	TotalSupply    float64 `json:"total_supply"`
	Platform       string
	Quote          map[string]USD
	MarketCap      string
	Volume24       string
	Change1        string
	Red            bool
}

type CMCHolder struct {
	Data []CMC `json:"data"`
}
type USD struct {
	Price    float64
	Volume24 float64 `json:"volume_24h"`
	Change1  float64 `json:"percent_change_1h"`
	Cap      float64 `json:"market_cap`
}

func main() {
	pat := os.Getenv("CMC")
	jsonString := DoGet(pat, "v1/cryptocurrency/listings/latest?limit=1000")
	var cmcHolder CMCHolder
	json.Unmarshal([]byte(jsonString), &cmcHolder)

	r := Replacer{}
	r.Caps = map[string]string{}
	r.Vol24 = map[string]string{}
	r.Change1 = map[string]string{}
	list := []CMC{}
	valid := `bitcoin ethereum binance-coin cardano dogecoin polkadot-new polygon stellar vechain terra-luna iota kusama tezos cosmos avalanche algorand elrond-egld qtum harmony nano celo kava`
	validMap := map[string]bool{}
	for _, v := range strings.Split(valid, " ") {
		validMap[v] = true
	}
	for _, c := range cmcHolder.Data {
		if validMap[c.Slug] == false {
			continue
		}
		fmt.Println(c.Symbol, c.Name, c.Slug)
		c.DateAdded = c.DateAdded[0:4]
		usd := c.Quote["USD"].Price
		mcap := (c.Circulating / 1000000000.0) * usd

		c.MarketCap = fmt.Sprintf("%0.2f", mcap)
		c.Volume24 = fmt.Sprintf("%0.2f", c.Quote["USD"].Volume24/1000000000.0)
		c.Change1 = fmt.Sprintf("%0.2f", c.Quote["USD"].Change1)
		if c.Symbol == "ALGO" || c.Symbol == "ADA" {
			c.Red = true
		}
		list = append(list, c)
	}
	r.Things = list
	MakeHtml(r)
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
