package main

import (
	"time"
)

func _2main() {
	ch1 := make(chan int)

	go func() {
		for {
			c := time.After(time.Second * 3)
			<-c
			println(2)
		}
	}()
	<-ch1

}
