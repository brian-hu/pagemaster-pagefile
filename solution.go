package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
    b, err := mem.ReadAddress(0)
    if err != nil {
        return nil, err
    }
    fmt.Println(b)
    b, err = mem.ReadAddress(12312442)
    if err != nil {
        return nil, err
    }
    fmt.Println(b)

    bs, err := mem.ReadPage(0)
    if err != nil {
        return nil, err
    }
    fmt.Println(string(bs))

    bs, err := mem.ReadPage(1000000)
    if err != nil {
        return nil, err
    }
    fmt.Println(string(bs))
	return "", nil
}
