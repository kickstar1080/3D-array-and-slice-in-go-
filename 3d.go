package main

import "fmt"

func main() {
	fmt.Printf("hello world")
	//for array we can build it like this
	var array = [2][3][3]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18},
		},
	}
	//printing
	fmt.Printf("\n = ", array)

	//accessing the specific element
	fmt.Printf("\n9th element = %d\n", array[0][2][2])

	//creating 3D slices
	var n int
	var m int
	var p int
	fmt.Printf("\n enter the value of n m p\n")
	//input
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	fmt.Scanln(&p)

	slice1 := make([][][]int, n)

	for i := 0; i < n; i++ {
		slice1[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			slice1[i][j] = make([]int, p)
		}
	}
	//test print
	fmt.Printf("\n %d\n", slice1)

}
