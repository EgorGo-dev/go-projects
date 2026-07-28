package main

import (
	"fmt"
)

func new() {
	fmt.Println("Ok let`s Go")

	arr := []int{1, 2, 3, 4, 5}

	i := 3
	j := 1
	fmt.Println(arr[i:])
	fmt.Println(arr[:j])
}
