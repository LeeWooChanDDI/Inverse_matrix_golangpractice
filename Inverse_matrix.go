package main

import "fmt"

func main() {
	var inverse_Matrix = [2][2]int{
		{7, 3},
		{5, 2},
	}

	var exist bool

	d := inverse_Matrix[0][0]*inverse_Matrix[1][1] - inverse_Matrix[0][1]*inverse_Matrix[1][0]

	if d != 0 {
		var A = [2][2]int{
			{inverse_Matrix[1][1], inverse_Matrix[0][1]},
			{inverse_Matrix[1][0], inverse_Matrix[0][0]},
		}

		exist = true
		fmt.Println(exist)
		fmt.Println(-A[0][0], A[0][1])
		fmt.Println(A[1][0], -A[1][1])
	} else {
		exist = false

		fmt.Println(exist)
	}
}
