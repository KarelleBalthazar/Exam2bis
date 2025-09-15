package main

import (
	"fmt"
)

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}

func main() {
	words := []string{"Bonjour", "roi", "bob"}

	isTwoLetters := func(s string) bool {
		return len(s) == 2
	}

	fmt.Println(CountIf(isTwoLetters, words))

}
