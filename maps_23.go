package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	a := strings.Split(s, " ")

	m := make(map[string]int)

	for _, v := range a {
		m[v] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
