package main

import "fmt"

func main() {
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m5 := []int{}
	fmt.Println(list)
	add(list, &m5)
	fmt.Println("Os números dessa lista maiores que 5 são :")
	fmt.Println(m5)
}
func add(list [10]int, m5 *[]int) {
	for _, i := range list {
		if i > 5 {
			*m5 = append(*m5, i)
		}
	}
}
