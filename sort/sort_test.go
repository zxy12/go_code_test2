package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	defer func() {
		r := recover()
		fmt.Println("recovery")
		fmt.Println(r)
	}()
	fmt.Println("111")
	ret := m.Run()
	os.Exit(ret)
}

func TestTTTT(m *testing.T) {
	fmt.Println("xxxx")
}
func TestTTTTb(m *testing.T) {
	fmt.Println("bbbxxxx")
}
