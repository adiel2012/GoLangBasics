package main

import "fmt"

func MainSlices() {
	fmt.Println("==============================================================")
	fmt.Println("Slices")
	fmt.Println("==============================================================")

	var arr1 []int = []int{781, 22, 43, 94, 345, 346, 987, 8, 9}

	slice1 := arr1[2:5]

	for _, v := range slice1 {
		fmt.Println(v)
	}

	mat1 := [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}

	slice2 := mat1[1:3]
	fmt.Println(slice2)

}
