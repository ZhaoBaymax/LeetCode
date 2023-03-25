package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	low_map := map[uint8]string{
		'a': "2",
		'b': "2",
		'c': "2",
		'd': "3",
		'e': "3",
		'f': "3",
		'g': "4",
		'h': "4",
		'i': "4",
		'j': "5",
		'k': "5",
		'l': "5",
		'm': "6",
		'n': "6",
		'o': "6",
		'p': "7",
		'q': "7",
		'r': "7",
		's': "7",
		't': "8",
		'u': "8",
		'v': "8",
		'w': "9",
		'x': "9",
		'y': "9",
		'z': "9",
	}
	up_map := map[uint8]string{
		'Z': "a",
		'A': "b",
		'B': "c",
		'C': "d",
		'D': "e",
		'E': "f",
		'F': "g",
		'G': "h",
		'H': "i",
		'I': "j",
		'J': "k",
		'K': "l",
		'L': "m",
		'M': "n",
		'N': "o",
		'O': "p",
		'P': "q",
		'Q': "r",
		'R': "s",
		'S': "t",
		'T': "u",
		'U': "v",
		'V': "w",
		'W': "x",
		'X': "y",
		'Y': "z",
	}
	ans := ""
	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' {
			ans += low_map[str[i]]
		} else if 'A' <= str[i] && str[i] <= 'Z' {
			ans += up_map[str[i]]
		} else {
			ans += string(str[i])
		}
	}
	fmt.Println(ans)
}
