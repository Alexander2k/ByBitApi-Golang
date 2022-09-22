package main

import (
	"ByBitApiV2/internal/client"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := client.NewAPI()
	tickerS, _ := a.Spot.GetTicker()
	tickerM, _ := a.Margin.GetTicker()

	fmt.Printf("%+v", tickerS)
	fmt.Printf("%+v", tickerM)
}
