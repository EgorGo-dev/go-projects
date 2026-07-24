package main

import (
    "fmt"
    "io"
    "net/http"
    "sort"
    "strings"
)

// toInt преобразует строку с числом в int (ручной аналог strconv.Atoi)
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

// BestStudents возвращает имена студентов с оценками выше среднего
func BestStudents(names []string) (string, error) {
    // карта: имя → оценка
    marks := make(map[string]int)
    sum := 0
    count := 0

    for _, name := range names {
        markStr, err := getMark(name)
        if err != nil {
            return "", err
        }

        mark := toInt(markStr)
        marks[name] = mark
        sum += mark
        count++
    }

    if count == 0 {
        return "", fmt.Errorf("нет студентов для подсчёта")
    }

    average := sum / count

    // отбираем студентов с оценкой выше среднего
    var best []string
    for name, mark := range marks {
        if mark > average {
            best = append(best, name)
        }
    }

    sort.Strings(best)                 // сортируем по алфавиту
    result := strings.Join(best, ",") // склеиваем через запятую

    return result, nil
}