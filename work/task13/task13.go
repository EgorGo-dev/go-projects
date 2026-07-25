package main

import (
    "errors"
    "time"
)

// fib рекурсивно вычисляет число Фибоначчи (без мемоизации)
func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

// TimeoutFibonacci вычисляет N-е число Фибоначчи с ограничением по времени
func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
    // Отрицательные числа недопустимы
    if n < 0 {
        return 0, errors.New("n must be non-negative")
    }

    ch := make(chan int)

    // Вычисление в фоновой горутине
    go func() {
        ch <- fib(n)
    }()

    // Ожидаем результат или таймаут
    select {
    case res := <-ch:
        return res, nil
    case <-time.After(timeout):
        return 0, errors.New("timeout")
    }
}