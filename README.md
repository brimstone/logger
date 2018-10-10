logger
======

My first attempt at a logger package with the usability I want.

Usage
-----
```
package main

import (
	"time"

	"github.com/brimstone/logger"
)

func main() {
	log := logger.Method("main")
	defer log.Profile(time.Now())

	log.Debug("A walrus appears",
		log.Field("animal", time.Now()),
	)

	log.Println("lol, so")
	time.Sleep(time.Second)
	log.Info("Just an info message, like, got this far")
}
```
