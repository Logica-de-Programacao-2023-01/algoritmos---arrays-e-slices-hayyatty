package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)
	list = append(list[:3], list[4:]...)
	fmt.Println(list)
}
