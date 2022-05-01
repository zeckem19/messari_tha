package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Expected output
type Result struct {
	Market                        int     `json:"market"`
	Total_volume                  float64 `json:"total_volume"`
	Mean_price                    float64 `json:"mean_price"`
	Mean_volume                   float64 `json:"mean_volume"`
	Volume_weighted_average_price float64 `json:"volume_weighted_average_price"`
	Percentage_buy                float32 `json:"percentage_buy"`
}

type Store struct {
	transactions        int
	volume, total_price float64
	buy                 int
}

type Trade struct {
	Id     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	Is_buy bool    `json:"is_buy"`
}

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

	for i, s := range record {
		if (Store{} != s) {
			r, err := json.Marshal(Result{
				Market:                        i,
				Total_volume:                  s.volume,
				Mean_price:                    s.total_price / float64(s.transactions),
				Mean_volume:                   s.volume / float64(s.transactions),
				Volume_weighted_average_price: s.total_price / s.volume,
				Percentage_buy:                float32(s.buy) / float32(s.transactions),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(string(r))
				fmt.Println(s.transactions)
			}
		}
	}
}

func process(t Trade) {
	m := t.Market
	record[m].transactions += 1
	record[m].volume += t.Volume
	record[m].total_price += (t.Price * t.Volume)
	// record[m].mean_price = record[m].total_price / record[m].volume
	// record[m].mean_volume = record[m].volume / float64(record[m].transactions)

	if t.Is_buy {
		record[m].buy += 1
	}
}
