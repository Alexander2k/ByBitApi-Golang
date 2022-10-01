package main

import (
	"ByBitApiV2/internal/client"
	"ByBitApiV2/internal/derivatives"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := client.NewByBitAPI()
	//start := Date(2022, 1, 1)
	//end := Date(2022, 2, 1)

	riskLimit, _ := api.Derivatives.GetOpenInterest(derivatives.OpenInterest{
		Category: "linear",
		Symbol:   "BTCUSDT",
		Interval: "5min",
	})

	fmt.Printf("%+v\n", riskLimit.Result.List)

}
