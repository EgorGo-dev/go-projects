package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan struct{})
  // горутина, которая асинхронно производит вычисления
  go func() {
    fmt.Println("начинаем вычисления...")
    // имитируем длинные вычисления
    time.Sleep(time.Second)
    fmt.Println("заканчиваем вычисления ...")
    // закрываем канал, чтобы получить сообщения
    close(ch)
  }()

  // программа блокируется
  <-ch
  fmt.Println("завершаем программу")
}
    