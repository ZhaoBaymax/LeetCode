package main

import "strings"

func interpret(command string) string {
	res := strings.Builder{}
	for i, v := range command {
		if v == 'G' {
			res.WriteByte('G')
			continue
		} else if v == '(' {
			if command[i+1] == ')' {
				res.WriteByte('o')
			} else {
				res.WriteString("a1")
			}
		}
	}
	return res.String()
}
