package main

import (
	"ByBitApiV2/internal/client"
	"ByBitApiV2/internal/derivatives"
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

	inst := a.Derivatives.GetInstrumentInfo(derivatives.INVERSE, "", "2")
	fmt.Printf("%+v\n", inst)
	fmt.Println("**************************************")

}
