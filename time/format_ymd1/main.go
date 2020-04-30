package main

import (
	"fmt"
	"time"
)

func TimeFormat(str string) int64 {
	formatTime, err := time.Parse("2006/1/02 15:04", str)
	if err != nil {
		fmt.Printf("err=%v", err)
		return 0
	}

	fmt.Println(formatTime)
	fmt.Println(formatTime.Unix())
	return formatTime.Unix()
}

func main() {
	str1 := "2020/1/22 3:40"
	TimeFormat(str1)
	str2 := "2020/11/22 3:40"
	TimeFormat(str2)
}
