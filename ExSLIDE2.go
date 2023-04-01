package main

import "fmt"

func main() {
	var (
		nome1, nome2, nome3 string
	)
	nomes := []string{}
	fmt.Println(nomes)
	fmt.Println("Digfite um nome")
	fmt.Scan(&nome1)
	nomes = append(nomes, nome1)
	fmt.Println(nomes)
	fmt.Println("Digite")
	fmt.Scan(&nome2)
	fmt.Println("Digite")
	fmt.Scan(&nome3)
	nomes = append(nomes, nome2, nome3)
	fmt.Println(nomes)

}
