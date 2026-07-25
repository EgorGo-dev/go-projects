package main

import (
    "time"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
    // Таймер на 0.1 секунды — закрывает stop-канал
    time.AfterFunc(100*time.Millisecond, func() {
        close(stop)
    })

    for i := 2; i <= N; i++ {
        // Проверяем сигнал остановки
        select {
        case <-stop:
            close(prime_nums) // закрываем канал перед выходом
            return
        default:
        }

        if isPrime(i) {
            // Проверяем остановку и при отправке
            select {
            case prime_nums <- i:
            case <-stop:
                close(prime_nums)
                return
            }
        }
    }

    // Если дошли до конца — закрываем канал
    close(prime_nums)
}