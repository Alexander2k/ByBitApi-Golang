package main

import (
	"ByBitApiV2/internal/client"
	"ByBitApiV2/internal/derivatives"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strconv"
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

	api := client.NewAPI()
	start := Date(2022, 1, 1)
	end := Date(2022, 2, 1)

	inst, _ := api.Derivatives.GetIndexPriceKline(derivatives.KlineParams{
		Category: "linear",
		Symbol:   "BTCUSDT",
		Interval: "5",
		Start:    strconv.FormatInt(start.UnixMilli(), 10),
		End:      strconv.FormatInt(end.UnixMilli(), 10),
		Limit:    "25",
	})
	fmt.Printf("%+v\n", inst)

}
