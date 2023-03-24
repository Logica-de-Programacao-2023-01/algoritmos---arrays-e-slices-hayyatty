package main

import (
	"fmt"
)

func main() {
	var mult float32
	mult = 1
	list := [4]float32{3.5, 6.4, 9.1, 2.7}
	fmt.Println(list)
	for _, num := range list {
		mult *= num
	}
	fmt.Println(mult)
}
