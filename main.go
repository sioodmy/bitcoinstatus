package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/sioodmy/bitcoinstatus/internal/price"
	"github.com/sioodmy/bitcoinstatus/internal/status"
)

func main() {

	min := 700
	max := 4000

	log.Println("Starting")

	for {
		rng := rand.Intn(max-min+1) + min
		interval := time.Duration(rng) * time.Second

		log.Println("Sleeping for ", interval)
		time.Sleep(interval)

		log.Println("Updating status at ", time.Now())
		updateStatus()
	}

}

func updateStatus() {
	price, err := price.GetPrice()
	if err != nil {
		log.Println("Couldn't set status: ", err)

	}
	status.SetStatus(price)

}
