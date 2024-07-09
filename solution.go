package main

import (
	"fmt"
)

func SearchString(mem *Memory) (string, error) {
	memAddress := 0
	var window []byte
	for {
		b, err := mem.ReadAddress(memAddress)
		if err != nil {
			return "", err
		}
		fmt.Print(string(b))

		if len(window) < 5 {
			window = append(window, b)
		} else {
			if string(window) == "gc24{" {
				fmt.Println("here")
				break
			} else {
				window = append(window, b)[1:]
			}
		}

		memAddress++
	}

	return "", nil
}
