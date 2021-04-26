package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	tablewriter "github.com/olekukonko/tablewriter"
	coingecko "github.com/superoo7/go-gecko/v3"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Xmr      float32 `yaml:"xmr"`
	Eth      float32 `yaml:"eth"`
	Doge     float32 `yaml:"doge"`
	Xtz      float32 `yaml:"xtz"`
	Nucypher float32 `yaml:"nucypher"`
}

func (balance *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("balance.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, balance)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return balance
}

func main() {
	var balance conf
	balance.getConf()
	p := message.NewPrinter(message.MatchLanguage("en"))

	cg := coingecko.NewClient(nil)
	coins := []string{"monero", "ethereum", "dogecoin", "tezos", "nucypher"}
	fiats := []string{"usd", "eur"}
	dt := time.Now()
	prices, err := cg.SimplePrice(coins, fiats)
	if err != nil {
		log.Fatal(err)
	}
	xmr := (*prices)["monero"]
	eth := (*prices)["ethereum"]
	doge := (*prices)["dogecoin"]
	tezos := (*prices)["tezos"]
	nucypher := (*prices)["nucypher"]

	totEur := xmr["eur"]*balance.Xmr + eth["eur"]*balance.Eth + doge["eur"]*balance.Doge + tezos["eur"]*balance.Xtz + nucypher["eur"]*balance.Nucypher
	totUsd := xmr["usd"]*balance.Xmr + eth["usd"]*balance.Eth + doge["usd"]*balance.Doge + tezos["usd"]*balance.Xtz + nucypher["usd"]*balance.Nucypher

	data := [][]string{
		[]string{
			"XMR",
			p.Sprintf("%f", balance.Xmr),
			p.Sprintf("%.4f", xmr["usd"]),
			p.Sprintf("%.2f", xmr["usd"]*balance.Xmr),
			p.Sprintf("%.4f", xmr["eur"]),
			p.Sprintf("%.2f", xmr["eur"]*balance.Xmr),
		},
		[]string{
			"ETH",
			p.Sprintf("%f", balance.Eth),
			p.Sprintf("%.4f", eth["usd"]),
			p.Sprintf("%.2f", eth["usd"]*balance.Eth),
			p.Sprintf("%.4f", eth["eur"]),
			p.Sprintf("%.2f", eth["eur"]*balance.Eth),
		},
		[]string{
			"DOGE",
			p.Sprintf("%f", balance.Doge),
			p.Sprintf("%.4f", doge["usd"]),
			p.Sprintf("%.2f", doge["usd"]*balance.Doge),
			p.Sprintf("%.4f", doge["eur"]),
			p.Sprintf("%.2f", doge["eur"]*balance.Doge),
		},
		[]string{
			"XTZ",
			p.Sprintf("%f", balance.Xtz),
			p.Sprintf("%.4f", tezos["usd"]),
			p.Sprintf("%.2f", tezos["usd"]*balance.Xtz),
			p.Sprintf("%.4f", tezos["eur"]),
			p.Sprintf("%.2f", tezos["eur"]*balance.Xtz),
		},
		[]string{
			"NuCy",
			p.Sprintf("%f", balance.Nucypher),
			p.Sprintf("%.4f", nucypher["usd"]),
			p.Sprintf("%.2f", nucypher["usd"]*balance.Nucypher),
			p.Sprintf("%.4f", nucypher["eur"]),
			p.Sprintf("%.2f", nucypher["eur"]*balance.Nucypher),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Coin",
		"Amount",
		"Market 💵",
		"Balance 💵",
		"Market 💶",
		"Balance 💶",
	})
	table.SetFooter([]string{
		"🗓️",
		dt.Local().Format("2 Jan 2006 15:04"),
		"Total 💵",
		p.Sprintf("%.2f", totUsd),
		"Total 💶",
		p.Sprintf("%.2f", totEur),
	})
	table.SetBorder(true)
	table.AppendBulk(data)
	table.Render()
}
