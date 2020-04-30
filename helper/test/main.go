package main

import (
	"fmt"

	"github.com/zxy12/go_code_test/helper"
)

const UINT_MIN uint = 0

const UINT_MAX uint = ^UINT_MIN

const INT_MAX int = int(^uint(0) >> 1)

const INT_MIN int = ^INT_MAX

const UINT64_MIN uint64 = 0

const UINT64_MAX uint64 = ^UINT64_MIN

const INT64_MAX int64 = int64(^uint64(0) >> 1)

const INT64_MIN int64 = ^INT64_MAX

func main() {

	fmt.Printf("uint max=[%v],uint min=[%v],int max[%v], int min=[%v]\n,", UINT_MAX, UINT_MIN, INT_MAX, INT_MIN)
	fmt.Printf("uint64 max=[%v],uint64 min=[%v],int64 max[%v], int64 min=[%v]\n,", UINT64_MAX, UINT64_MIN, INT64_MAX, INT64_MIN)

	str := "18446744073709551615"
	var c helper.ComType
	err := (&c).Load(str)
	fmt.Printf("com=%+v\n", c)
	if err != nil {
		fmt.Println("load fail:", err)
		return
	}

	s, err := c.String()
	fmt.Println(s, err)

	i, err := c.Int()
	fmt.Println(i, err)

	ui, err := c.Uint()
	fmt.Println(ui, err)

	i64, err := c.Int64()
	fmt.Println(i64, err)

	ui64, err := c.Uint64()
	fmt.Println(ui64, err)

	p := helper.NewPrinter()
	p2 := helper.NewPrinter()

	p.Trace("p1-1111111111")
	p.Trace("p1-2222222222")
	p2.Trace("p2-1111111111")
	p.Trace("p1-3333333333")
	p2.Trace("p2-222222222")

}
