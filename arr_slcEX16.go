package main

import "fmt"

func main() {
	var ()
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	div2 := []int{}
	fmt.Println(list)
	for _, i := range list {
		if i%2 == 0 {
			div2 = append(div2, i)
		}
	}
	fmt.Println("Os números dessa lista divisies por 2 são :")
	fmt.Println(div2)
}
