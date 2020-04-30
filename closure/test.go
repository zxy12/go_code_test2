package main

import "fmt"

func a() int {
	a := 0
	f := func() int {
		a++
		return a
	}
	return f()
}

func b() func() int {
	a := 0
	f := func() int {
		a++
		return a
	}
	return f
}

func main() {
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	f := b()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
