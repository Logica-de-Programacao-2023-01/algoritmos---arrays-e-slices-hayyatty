package main

import "fmt"

func main() {
	var (
		n1, nult int
	)
	list := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(list)
	fmt.Println("Qual número deve ser inserido em primeiro ?")
	fmt.Scan(&n1)
	fmt.Println("Qual número deve ser adicionado em último ?")
	fmt.Scan(&nult)
	list[0] = n1
	list[6] = nult
	fmt.Println(list)
}
