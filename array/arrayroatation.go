package main

import "fmt"

func Rotation() {
	array := []int{1, 2, 4, 5, 6, 7}

	result := make([]int, len(array))

	copy(result, array)

	fmt.Println(result)

}