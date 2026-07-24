package main

import (
    "fmt"
    "io"
    "net/http"
)

// getMark отправляет запрос к серверу и возвращает оценку студента в виде строки
func getMark(name string) (string, error) {
    // Собираем URL с именем студента
    // Например: http://localhost:8082/mark?name=Alice
    url := fmt.Sprintf("http://localhost:8082/mark?name=%s", name)

    // Отправляем GET-запрос
    resp, err := http.Get(url)
    if err != nil {
        // Если запрос не удался (нет сети, сервер не отвечает и т.д.)
        return "", fmt.Errorf("ошибка запроса: %w", err)
    }
    // Обязательно закрываем тело ответа после чтения
    defer resp.Body.Close()

    // Проверяем код ответа
    if resp.StatusCode == 404 {
        return "", fmt.Errorf("студент %s не найден", name)
    }
    if resp.StatusCode == 500 {
        return "", fmt.Errorf("ошибка сервера при получении оценки %s", name)
    }
    if resp.StatusCode != 200 {
        return "", fmt.Errorf("неизвестный код ответа: %d", resp.StatusCode)
    }

    // Читаем тело ответа (там должно быть число)
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("не удалось прочитать ответ: %w", err)
    }

    // Возвращаем оценку как строку
    return string(body), nil
}

// padMark дополняет строку нулями слева до длины 3
func padMark(s string) string {
    for len(s) < 3 {
        s = "0" + s
    }
    return s
}

// Compare сравнивает оценки двух студентов
func Compare(name1, name2 string) (string, error) {
    // Получаем оценки
    mark1, err := getMark(name1)
    if err != nil {
        return "", err
    }
    mark2, err := getMark(name2)
    if err != nil {
        return "", err
    }

    // Дополняем нулями, чтобы сравнивать строки как числа
    m1 := padMark(mark1)
    m2 := padMark(mark2)

    // Сравниваем
    if m1 > m2 {
        return ">", nil
    } else if m1 < m2 {
        return "<", nil
    }
    return "=", nil
}