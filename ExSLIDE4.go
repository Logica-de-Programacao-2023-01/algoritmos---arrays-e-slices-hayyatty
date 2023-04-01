package main

import "fmt"

func main() {
	var (
		n1, n2, n3, n4, n5, n6 float32
	)
	list := [6]float32{5.6, 6.4, 3.4, 7.9, 1.9, 4.7}
	fmt.Println(list)
	n1 = list[0]
	n2 = list[1]
	n3 = list[2]
	n4 = list[3]
	n5 = list[4]
	n6 = list[5]
	media := (n1 + n2 + n3 + n4 + n5 + n6) / 6
	fmt.Println(media)
}
