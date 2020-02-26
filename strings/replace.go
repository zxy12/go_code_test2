package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "fsf 131 \n \\n 1 "

	fmt.Println(s)
	fmt.Println(strings.Replace(s, "\n", " ", -1))

}
