package main

import( "strings"
        "strconv"
)

func HexToDecimal(word string) string {
	words := strings.Fields(word)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hex, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(hex, 10)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, " ")
}

func BinToDecimal(word string) string {
	words := strings.Fields(word)

	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			bin, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(bin, 10)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, " ")
}
