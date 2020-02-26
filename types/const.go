package main

import "fmt"

type Tint int

const (
	AA Tint = 1
	BB      = 2
	CC      = 3
)

const (
	_ Tint = iota
	AAA
	BBB
	CCC
)

func main() {
	fmt.Printf("%+v, %t, %+v, %t\n", AA, AA, BB, BB)
	fmt.Printf("%+v, %t, %+v, %t\n", AAA, AAA, BBB, BBB)
	var a int = 1

	switch AAA {
	case Tint(a):
		println("AAA=1")
	}
}
