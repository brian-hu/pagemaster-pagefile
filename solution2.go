package main

import (
	"fmt"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func SearchString(mem *Memory) (string, error) {
	bs, err := mem.ReadPage(0)
	if err != nil {
		return "", err
	}

    decoder := unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
    converted, err := decoder.Bytes(bs)
    if err != nil {
        return "", nil
    }
    fmt.Println(string(converted))
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
