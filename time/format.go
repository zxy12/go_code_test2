package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	v := time.Now().Format("060102150405")
	fmt.Println(v)
	rand.Seed(time.Now().UnixNano())
	r := fmt.Sprintf("%03v", rand.Intn(1000))
	fmt.Println(r)



}
