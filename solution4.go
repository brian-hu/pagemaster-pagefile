package main

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
                out := []byte("gc24{")

				for {
					char, err := mem.ReadAddress(memAddress)
					if err != nil {
						return "", err
					}
                    out = append(out, char)

					if string(char) == "}" {
						return string(out), nil
					}
                    memAddress++
				}

			} else {
				window = append(window, b)[1:]
			}
		}

		memAddress++
	}
}
