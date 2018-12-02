package main

import (
	"github.com/brimstone/logger"
)

func main() {
	/*
		log := logger.New(&logger.Options{
			Method: "main",
			Delay:  time.Second * 1,
		})
	*/
	log := logger.Method("main")
	log.Println("Starting")

	for count := 0; count < 100000000; count++ {
		log.Counter("count", 1)
	}
	log.Println("End")
}
