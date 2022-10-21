package main

import (
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	sentence_list := strings.Split(sentence, " ")
	for i, v := range sentence_list {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
	}
	return -1
}
