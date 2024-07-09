package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
    fmt.Println(mem.ReadAddress(0))
	fmt.Println("here")
	return "asdf", nil
}
