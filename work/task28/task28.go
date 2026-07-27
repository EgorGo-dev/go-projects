package main

import (
    "strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
    var parts []string
    parts = append(parts, "curl")

    // Для GET не добавляем -X (по умолчанию curl использует GET)
    if method != "GET" {
        parts = append(parts, "-X", method)
    }

    if headers != "" {
        // Разбиваем заголовки по \n
        headerLines := strings.Split(headers, "\n")
        for _, h := range headerLines {
            h = strings.TrimSpace(h)
            if h == "" {
                continue
            }
            parts = append(parts, "-H", "'"+h+"'")
        }
    }

    if body != "" {
        parts = append(parts, "--data", "'"+body+"'")
    }

    parts = append(parts, url)

    // Склеиваем все части через пробел
    return strings.Join(parts, " ")
}