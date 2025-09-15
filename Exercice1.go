package main

func AppendRange(min, max int) []int {
	var tab []int
	for i := min; i <= max; i++ {
		tab = append(tab, i)
	}
	return tab
}
