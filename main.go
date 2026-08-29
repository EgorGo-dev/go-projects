package main 

import (
    "fmt"
)

func Sum(num1, num2 int) string {
    return fmt.Sprintf("Сумма чисел %d и %d: %d", num1, num2, num1 + num2)
}

func main() {
    var a, b int 
    fmt.Print("Введите 2 числа (через пробел): ")
    _, err := fmt.Scan(&a, &b)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
        return
    }
    fmt.Println(Sum(a, b))
}