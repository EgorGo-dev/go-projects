package main

import (
    "context"
    "os"
)

// readJSON читает файл по пути path и отправляет его содержимое в канал result.
// Учитывает отмену контекста.
func readJSON(ctx context.Context, path string, result chan<- []byte) {
    defer close(result)

    // Проверяем контекст до чтения файла
    select {
    case <-ctx.Done():
        return
    default:
    }

    // Читаем файл
    data, err := os.ReadFile(path)
    if err != nil {
        return
    }

    // Проверяем контекст перед отправкой
    select {
    case <-ctx.Done():
        return
    default:
    }

    // Отправляем данные в канал
    select {
    case result <- data:
    case <-ctx.Done():
        return
    }
}