package CollectCandle

import (
	"fmt"
	"github.com/Alexander2k/ByBitApi-Golang/config"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/client"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/derivatives"
	"log"
	"os"
	"strconv"
	"time"
)

func Date(year, month, day, hour, min int) int64 {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).UnixMilli()
}

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	api := client.NewByBitAPI(conf)

	offset := int64(12000000)
	symbol := "XRPUSDT"

	start := Date(2021, 5, 13, 0, 0)
	end := Date(2023, 7, 13, 0, 0)

	file, err := os.Create(symbol + ".txt")
	if err != nil {
		log.Printf("Error creating file")
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Printf("Error closing file: %v\n", err)
		}
	}(file)
	_, err = file.WriteString("<TICKER>,<PER>,<DATE>,<TIME>,<OPEN>,<HIGH>,<LOW>,<CLOSE>,<VOL>\n")
	if err != nil {
		log.Printf("Error writing: %v\n", err)
		return
	}

	for {

		kl := derivatives.KlineParams{
			Category: "linear",
			Symbol:   symbol,
			Interval: "1",
			Start:    strconv.FormatInt(start, 10),
			End:      strconv.FormatInt(start+offset, 10),
			Limit:    "",
		}

		data, err := api.Derivatives.GetKline(kl)
		if err != nil {
			log.Printf("Error getting kline data: %v\n", err)

		}

		for i := len(data.Result.List) - 1; i >= 0; i-- {
			c := data.Result.List[i]
			date, _ := strconv.ParseInt(c[0], 10, 64)
			open, high, low, cl, vol := c[1], c[2], c[3], c[4], c[5]

			s := fmt.Sprintf("%s,%d,%s,%s,%s,%s,%s,%s\n", data.Result.Symbol, 1, time.UnixMilli(date).Format("20060102,150405"), open, high, low, cl, vol)
			_, err := file.WriteString(s)
			if err != nil {
				log.Printf("Error writing: %v\n", err)
			}

		}

		start = start + offset
		//time.Sleep(time.Second)

		if start >= end {
			log.Println("Done writing")
			break
		}

	}
}
