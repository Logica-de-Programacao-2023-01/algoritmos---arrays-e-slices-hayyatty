package main

import "fmt"

func main() {
	list := []int{2, 3, 5, 6, 8, 9, 12, 14}
	mult3 := []int{}
	for _, i := range list {
		if i%3 == 0 {
			mult3 = append(mult3, i)
		}
	}
	fmt.Println(list)
	fmt.Println("Os multiplos de 3 nessa lista são:", mult3)
}
