package main

import "fmt"

func main() {
	var (
		n     int
		found bool = false
	)
	list := []int{1, 2, 1, 3, 4, 5}
	fmt.Println(list)
	fmt.Println("Digite um número para remover")
	fmt.Scan(&n)

	for i := 0; i < len(list); i++ {
		if list[i] == n {
			list = append(list[:i], list[i+1:]...)
			found = true
		}
	}
	if found {
		fmt.Println("Numeros removido")
		fmt.Println(list)
	} else {
		fmt.Println("O número não esta na lista")
	}
}
