package main

import (
	"fmt"
)

func test(input []int) []int {
	// input: 0x40c110->{address, 4, 4}
	fmt.Printf("Slice inside address: %p\n", &input)
	// input: 0x40c110->{address, 4, 4}，但对address的修改影响到了val:0x40c0e0->{address, 4, 4}
	input[0] = -1
	input = append(input, 5)
	input[0] = -2
	fmt.Printf("Slice inside address2: %p\n", &input)
	// input: 0x40c110->{address, 5, 8}，但val仍然为0x40c0e0->{address, 4, 4}
	fmt.Println(input, len(input), cap(input))
	return input
}

func modify(ages []int) {
	fmt.Printf("函数里接收到slice的内存地址是%p\n", ages)
	ages[0] = 1
}

func main() {
	// val：0x40c0e0->{address, 4, 4}
	val := []int{1, 2, 3, 4}
	fmt.Printf("Slice outside address: %p\n", &val)
	fmt.Println(val, len(val), cap(val))
	copyVal := test(val)
	fmt.Printf("Slice outside address: %p\n", &val)
	fmt.Println(val, len(val), cap(val))
	fmt.Printf("CopySlice address: %p\n", &copyVal)
	fmt.Println(copyVal, len(copyVal), cap(copyVal))

	ages := []int{6, 6, 6}
	fmt.Printf("原始slice的内存地址是%p\n", ages)
	modify(ages)
	fmt.Println(ages)

}
