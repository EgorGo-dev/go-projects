package main 

import (
	"fmt"
	"math/rand"
)

func Random(n int) bool {
	return rand.Intn(10) == n
}

func main() {
	var n int
	fmt.Scan(&n)
	if ok := Random(n); !ok {
		fmt.Println("Ты не угадал попробуй в следующий раз")
	} else {
		fmt.Println("Ты угадал!!")
	}
}