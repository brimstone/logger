package main

import (
	"time"

	"github.com/brimstone/logger"
)

var log = logger.New()

func main() {
	/*
		log := logger.New(&logger.Options{
			Method: "main",
			Delay:  time.Second * 1,
		})
	*/
	defer log.Profile(time.Now())
	log.Println("Starting")

	for count := 0; count < 100000000; count++ {
		log.Counter("count", 1)
	}
	log.Println("End")
}
