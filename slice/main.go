package main

import "fmt"
import "reflect"

func main() {

	s := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	print(&s)

	s1 := s[0:]
	print(&s1)
	s1[0] = 99
	println("after change s1")
	print(&s)
	print(&s1)

	println()
	s1 = s[0:0]
	//print(&s1)
	s1 = append(s1, 999)
	println("after change s1")
	print(&s)
	print(&s1)

	println()
	s1 = s[0:1:1]
	print(&s1)
	s1 = append(s1, 99999)
	println("after change s1")
	print(&s)
	print(&s1)

	println("----")
	a := reflect.New(reflect.TypeOf(s))
	fmt.Printf("%+v\n", a.Interface())
	print(a.Interface())
	//a = append(a, 99999)
	//println("after change s1")
	//print(&s)
	//print(&a)

}

func print(s *[]int64) {
	fmt.Printf("s=%p: len=%v,cap=%v,first=%v,last=%v\n", s, len(*s), cap(*s), (*s)[0], (*s)[len(*s)-1])
}
