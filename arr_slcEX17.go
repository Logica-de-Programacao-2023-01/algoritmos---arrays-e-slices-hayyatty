package main

import (
	"fmt"
)

func main() {
	var (
		som int
	)
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(list); i++ {
		som += list[i]
	}
	for v, i := range list {
		if v%2 == 0 {
			list[i] = som
		}

	}
	fmt.Println(list)

}
