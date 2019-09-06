package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"sample"
)

func Find10Frequent(s string) []string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9\\s]+")

	processedString := reg.ReplaceAllString(s, "")
	words := strings.Split(processedString, " ")
	wordCounter := make(map[string]int)
	res := make([]string, 0, 10)

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		wordCounter[strings.ToLower(word)] += 1
	}

	keys := make([]string, 0, len(wordCounter))
	for k := range wordCounter {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return wordCounter[keys[i]] > wordCounter[keys[j]] })

	for i := 0; i < 10 && i < len(keys); i++ {
		res = append(res, keys[i])
	}

	return res
}

func main() {
	fmt.Println(Find10Frequent("Hello Man     hello !!!,:sdfasdf"))
}
