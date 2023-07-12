package main

import (
	"fmt"
	"github.com/Alexander2k/ByBitApi-Golang/config"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/client"
	"time"
)

func Date(year, month, day int) int64 {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).UnixMilli()
}
func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	api := client.NewByBitAPI(conf)
	//start := Date(2023, 1, 1)
	//end := Date(2023, 2, 1)
	//symbol := "BTCUSDT"
	//
	////kl := derivatives.KlineParams{
	////	Category: "linear",
	////	Symbol:   symbol,
	////	Interval: "1",
	////	Start:    strconv.FormatInt(start, 10),
	////	End:      strconv.FormatInt(end, 10),
	////	Limit:    "",
	////}
	data, _ := api.Spot.GetTicker("")
	for _, d := range data.Result.List {
		fmt.Printf("%+v\n", d)
	}
	//fmt.Printf("%+v", data)

}
