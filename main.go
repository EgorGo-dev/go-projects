package main 

import (
    "fmt"
)

func describe(input interface{}) string {
    switch v := input.(type) {
    case int:
        return fmt.Sprintf("Это число: %d", v)
    case string:
        return fmt.Sprintf("Это строка: %s", v)
    case bool:
        return fmt.Sprintf("Это булево значение: %t", v)
    default:
        return "Неизвестный тип"
    }
}

func main() {
    fmt.Println(describe(10))
    fmt.Println(describe("Егор"))
    fmt.Println(describe(true))
    fmt.Println(describe(3.14))
}