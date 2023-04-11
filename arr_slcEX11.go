package main

import "fmt"

func main() {
	var (
		l int
		c int
	)
	list := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Digite a linha do número que deseja:")
	fmt.Scan(&l)
	fmt.Println("Digite a coluna do número que deseja:")
	fmt.Scan(&c)
	if l >= 0 && l < len(list) && c >= 0 && c < len(list[0]) {
		fmt.Printf("O valor da posiçao [%d] e [%d]:%d\n", l, c, list[l][c])
	}
}
