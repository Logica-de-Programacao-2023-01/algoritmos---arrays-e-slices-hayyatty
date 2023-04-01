package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	var (
		n   int
		est int = 0
	)
	fmt.Println(list)
	fmt.Println("Digite um número para adicionar ")
	fmt.Scan(&n)
	for i := 0; i < len(list); i++ {
		if n == list[i] {
			est++
		} else {
			continue
		}

	}
	if est == 0 {
		list = append(list, n)
		fmt.Println(list)
	} else if est != 0 {
		fmt.Println("O numero já esta na lista")
	}

}
