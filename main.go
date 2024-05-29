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
	max := 1440

	log.Println("Starting")

	for {
		rng := rand.Intn(max-min+1) + min
		interval := time.Duration(rng) * time.Second

		updateStatus()
		log.Println("Updating status at ", time.Now())

		log.Println("Sleeping for ", interval)
		time.Sleep(interval)

	}

}

func updateStatus() {
	price, err := price.GetPrice()
	if err != nil {
		log.Println("Couldn't set status: ", err)

	}
	status.SetStatus(price)

}
