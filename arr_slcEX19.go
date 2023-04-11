package main

import "fmt"

func main() {
	list1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listS := []int{}

	for i := 0; i < len(list1); i++ {
		soma := list1[i] + list2[i]
		listS = append(listS, soma)
	}
	fmt.Println(listS)

}
