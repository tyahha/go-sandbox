package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func countStringInStrings(t string, ts []string) int {
	count := 0
	for _, e := range ts {
		if t == e {
			count++
		}
	}
	return count
}

func WordCount(s string) map[string]int {
	tokens := strings.Split(s, " ")
	checkedTokens := make([]string, len(tokens))
	ret := make(map[string]int)
	for i, token := range tokens {
		if countStringInStrings(token, checkedTokens) <= 0 {
			ret[token] = countStringInStrings(token, tokens[i:])
			checkedTokens[i] = token
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
