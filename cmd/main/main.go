package main

import (
	"fmt"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/client"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/derivatives"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"time"
)

func Date(year, month, day int) int64 {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).UnixMilli()
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := client.NewByBitAPI()
	start := Date(2023, 1, 1)
	end := Date(2023, 2, 1)
	symbol := "BTCUSDT"

	kl := derivatives.KlineParams{
		Category: "linear",
		Symbol:   symbol,
		Interval: "1",
		Start:    strconv.FormatInt(start, 10),
		End:      strconv.FormatInt(end, 10),
		Limit:    "",
	}
	kline, _ := api.Derivatives.GetKline(kl)
	fmt.Printf("%+v", kline)

}
