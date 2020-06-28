package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 8, 3}
	fmt.Println(len(a))
	//数组是值类型,不能直接排序，必须转为切片
	sort.Ints(a[:])
	fmt.Println(a)
}
func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	//按照字母顺序排序,从小到大
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 28.2, 392342.2, 0, 6}
	//从小到大排序
	sort.Float64s(a[:])
	fmt.Println(a)
}
func testIntSearch() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	//数组是值类型,不能直接排序，必须转为切片
	sort.Ints(a[:])
	//SearchInts默认排序后的位置，一定要排序后在查找
	index := sort.SearchInts(a[:], 348)

	fmt.Println(index)
}

func main() {
	testIntSort()
	testStrings()
	testFloat()
	testIntSearch()
}
