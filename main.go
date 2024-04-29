package main

import (
	"go-gateway/tb"
	"go-gateway/util"
	"time"
)

func main() {
	util.Load()

	tb.Connect()

	for {
		time.Sleep(10 * time.Second)
	}

}
