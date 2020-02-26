package main

import (
	"fmt"
	"time"
)

func _3main() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		time.Sleep(time.Second * 10)
		fmt.Println("Ticker running...")
	}
}
