package main

import (
	"fmt"
)

func main() {
	// ! ПРОВЕРКА КАК РАБОТАЕТ :
	fmt.Println("Ok let`s Go")
	arr := []int{1, 2, 3, 4, 5}

	i := 2
	j := 3
	fmt.Println(arr[i:])
	fmt.Println(arr[:j])
}