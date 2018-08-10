package main

import "fmt"

func main() {
	var a [5]int

	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("emp:", a[4])

	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i :=0; i<2; i++ {
		for j := 0; j<3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
