package main

import (
    "fmt"
    "io"
    "net/http"
)

// toInt преобразует строку с числом в int
func toInt(s string) int {
    result := 0
    for i := 0; i < len(s); i++ {
        digit := int(s[i] - '0')
        result = result*10 + digit
    }
    return result
}

// getMark запрашивает оценку студента у сервера
func getMark(name string) (string, error) {
    url := fmt.Sprintf("http://localhost:8082/mark?name=%s", name)

    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("ошибка запроса: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == 404 {
        return "", fmt.Errorf("студент %s не найден", name)
    }
    if resp.StatusCode == 500 {
        return "", fmt.Errorf("ошибка сервера при получении оценки %s", name)
    }
    if resp.StatusCode != 200 {
        return "", fmt.Errorf("неизвестный код ответа: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("не удалось прочитать ответ: %w", err)
    }

    return string(body), nil
}

// Average считает среднюю оценку по списку студентов
func Average(names []string) (int, error) {
    sum := 0
    count := 0

    for _, name := range names {
        markStr, err := getMark(name)
        if err != nil {
            return 0, err
        }

        mark := toInt(markStr)
        sum += mark
        count++
    }

    if count == 0 {
        return 0, fmt.Errorf("нет студентов для подсчёта среднего")
    }

    return sum / count, nil
}