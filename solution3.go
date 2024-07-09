package main

import (
	"strings"
)

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
			if string(window) == "gc24{" {
				var sb strings.Builder
				sb.WriteString("gc24{")

				for {
					char, err := mem.ReadAddress(memAddress)
					if err != nil {
						return "", err
					}
					sb.WriteString(char)

					if string(char) == "}" {
						return sb.String(), nil
					}
				}

			} else {
				window = append(window, b)[1:]
			}
		}

		memAddress++
	}
}
