package main

import (
	"fmt"
    "unicode/utf16"
)

func SearchString(mem *Memory) (string, error) {
	bs, err := mem.ReadPage(0)
	if err != nil {
		return "", err
	}
    fmt.Println(bs)

    shorts := make([]uint16, len(bs)/2)
	for i := 0; i < n; i += 2 {
		shorts[i/2] = (uint16(bs[i]) << 8) | uint16(bs[i+1])
	}

    utf8_runes := utf16.Decode(shorts)
    fmt.Println(string(utf8_runes))

	//memAddress := 0
	//var window []byte
	//for {
	//b, err := mem.ReadAddress(memAddress)
	//if err != nil {
	//return "", err
	//}
	//fmt.Print(string(b))

	//if len(window) < 5 {
	//window = append(window, b)
	//} else {
	//if string(window) == "grpc{" {
	//fmt.Println("here")
	//break
	//} else {
	//window = append(window, b)[1:]
	//}
	//}

	//memAddress += 8
	//}

	return "", nil
}
