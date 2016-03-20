package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("Strings are sorted:", sort.StringsAreSorted(strs))

	ints := []int{3, 1, 2}
	fmt.Println("Ints are sorted:", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println("ints", ints)
	fmt.Println("Ints are sorted:", sort.IntsAreSorted(ints))

}
