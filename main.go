package main

import (
	"fmt"
	"os"

	"github.com/obaydullahmhs/kafka-producer/pkg/producer"
)

func main() {
	fmt.Println("Starting...")
	producer, err := producer.NewProducer()
	fmt.Println(err)
	if err != nil {
		os.Exit(1)
	}
	producer.SendRandomMessages()
	fmt.Println("Done!")
}
