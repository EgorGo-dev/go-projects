package main

import (
    "bufio"
    "context"
    "encoding/json"
    "io"
    "strings"
    "time"
)

type Ticket struct {
    Ticket string    `json:"ticket"`
    User   string    `json:"user"`
    Status string    `json:"status"`
    Date   time.Time `json:"date"`
}

func GetTasks(ctx context.Context, r io.Reader, w io.Writer, user *string, status *string, timeout time.Duration) error {
    // 1. Создаём контекст с таймаутом
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    // 2. Результат будем собирать в слайс
    var result []Ticket

    // 3. Читаем строки из r
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        // Проверяем контекст (таймаут или отмена)
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }

        line := scanner.Text()
        if line == "" {
            continue
        }

        // 4. Парсим строку как раньше
        if !strings.HasPrefix(line, "TICKET-") {
            continue
        }

        parts := strings.SplitN(strings.TrimPrefix(line, "TICKET-"), "_", 4)
        if len(parts) != 4 {
            continue
        }

        ticketID := parts[0]
        userName := parts[1]
        ticketStatus := parts[2]
        dateStr := parts[3]

        // Проверяем статус
        if ticketStatus != "Готово" && ticketStatus != "В работе" && ticketStatus != "Не будет сделано" {
            continue
        }

        // Парсим дату
        date, err := time.Parse("2006-01-02", dateStr)
        if err != nil {
            continue
        }

        // Фильтр по пользователю
        if user != nil && *user != userName {
            continue
        }

        // Фильтр по статусу
        if status != nil && *status != ticketStatus {
            continue
        }

        // Добавляем в результат
        result = append(result, Ticket{
            Ticket: "TICKET-" + ticketID,
            User:   userName,
            Status: ticketStatus,
            Date:   date,
        })
    }

    // 5. Проверяем ошибки сканера
    if err := scanner.Err(); err != nil {
        return err
    }

    // 6. Записываем JSON в w
    encoder := json.NewEncoder(w)
    encoder.SetIndent("", "  ") // красивое форматирование (опционально)
    if err := encoder.Encode(result); err != nil {
        return err
    }

    return nil
}