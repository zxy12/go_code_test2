package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	fmt.Println("vim-go")
	sVals := map[int]S{1: {"A"}}

	// You can only call Read using a value
	sVals[1].Read()

	// This will not compile:
	//sVals[0].Write("test")

	sPtrs := map[int]*S{1: &S{"A"}}

	// You can call both Read and Write using a pointer
	sPtrs[1].Read()
	sPtrs[1].Write("test")

	fmt.Println(*sPtrs[1])

	println("start main")
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	/*
		for v := range ch {
			fmt.Println(v, ok)
		}
	*/
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v, ok)
		} else {
			fmt.Println(v, ok)
			break
		}
	}
	println("end main")

}
