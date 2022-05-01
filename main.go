package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PrettyData struct {
	Market                        int    `json:"market"`
	Total_volume                  string `json:"total_volume"`
	Mean_price                    string `json:"mean_price"`
	Mean_volume                   string `json:"mean_volume"`
	Percentage_buy                string `json:"percentage_buy"`
	Volume_weighted_average_price string `json:"volume_weighted_average_price"`
}

type Store struct {
	// Required for computation
	transactions int
	total_price  float64
	buy          int

	// Expected output
	market                        int
	total_volume                  float64
	mean_price                    float64
	mean_volume                   float64
	percentage_buy                float32
	volume_weighted_average_price float64
}

type Trade struct {
	Id     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	Is_buy bool    `json:"is_buy"`
}

// Define a global singleton
var record [13000]Store

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	t := Trade{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "BEGIN" {
			continue
		}
		if line == "END" {
			break
		}

		err := json.Unmarshal([]byte(line), &t)
		if err != nil {
			log.Fatal(err)
		} else {
			process(t)
		}
	}

	for _, s := range record {

		// Checks if market was traded, if it wasn't, s would be an empty struct
		if (Store{} != s) {
			pd := PrettyData{
				Market:                        s.market,
				Total_volume:                  fmt.Sprintf("%.2f", s.total_volume),
				Mean_price:                    fmt.Sprintf("%.5f", s.mean_price),
				Mean_volume:                   fmt.Sprintf("%.2f", s.mean_volume),
				Volume_weighted_average_price: fmt.Sprintf("%.2f", s.volume_weighted_average_price),
				Percentage_buy:                fmt.Sprintf("%.2f%%", s.percentage_buy*100),
			}
			r, err := json.Marshal(pd)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(string(r))
				// print number of transactions
				// fmt.Println(s.transactions)
			}
		}
	}
}

func process(t Trade) {

	record[t.Market].market = t.Market
	record[t.Market].transactions += 1
	record[t.Market].total_volume += t.Volume
	record[t.Market].total_price += (t.Price * t.Volume)
	record[t.Market].mean_price = record[t.Market].total_price / record[t.Market].total_volume
	record[t.Market].mean_volume = record[t.Market].total_volume / float64(record[t.Market].transactions)
	record[t.Market].volume_weighted_average_price = record[t.Market].total_price / record[t.Market].total_volume

	if t.Is_buy {
		record[t.Market].buy += 1
	}
	record[t.Market].percentage_buy = float32(record[t.Market].buy) / float32(record[t.Market].transactions)
}
