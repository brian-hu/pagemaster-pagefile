package main

import "fmt"
import "bytes"

func SearchString(mem *Memory) (string, error) {
    memAddress := 0
    var window []byte
    for {
        b, err := mem.ReadAddress(memAddress)
        if err != nil {
            return "", err
        }

        if len(window) < 5 {
            window = append(window, b)
        } else {
            if string(window) == "grpc{" {
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