package main

import (
	"fmt"
)

func IsSorted(cmp func(int, int) int, tab []int) bool {
	if len(tab) <= 1 {
		return true
	}
	for i := 1; i < len(tab); i++ {
		if cmp(tab[i-1], tab[i]) > 0 {
			return false
		}
	}
	return true
}
func compare(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(IsSorted(compare, []int{1, 2, 3, 4, 5}))
	fmt.Println(IsSorted(compare, []int{5, 4, 3, 2, 1}))
	fmt.Println(IsSorted(compare, []int{1, 3, 2, 4, 5}))
	fmt.Println(IsSorted(compare, []int{7}))
	fmt.Println(IsSorted(compare, []int{}))
}
