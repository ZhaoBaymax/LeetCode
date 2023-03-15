package main

import (
	"fmt"
	"sort"
	"strconv"
)

func topKFrequent(words []string, k int) []string {
	word_map := map[string]int{}
	for _, w := range words {
		word_map[w]++
	}
	word := make([]string, 0)
	for k, _ := range word_map {
		word = append(word, k)
	}
	sort.Slice(word, func(i, j int) bool {
		i_w, j_w := word[i], word[j]
		if word_map[i_w] == word_map[j_w] {
			return i_w < j_w
		}
		return word_map[i_w] > word_map[j_w]
	})
	return word[:k]
}
func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2))
	a, _ := strconv.Atoi("788")
	b, _ := strconv.Atoi("89")
	fmt.Println(a < b)
}
