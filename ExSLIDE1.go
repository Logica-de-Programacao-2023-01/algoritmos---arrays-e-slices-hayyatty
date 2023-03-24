package main

import "fmt"

func main() {
	list := [5]int{1, 2, 3, 4, 5}
	/*
		for _, list := range list {
			fmt.Println(list)
		}
	*/
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
