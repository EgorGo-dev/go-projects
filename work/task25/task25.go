package main

func BuildHTTPRequest(method, url, host, headers, body string) string {
    // Собираем первую строку: метод, путь и версия HTTP
    request := method + " " + url + " HTTP/1.1\r\n"

    // Добавляем Host
    request += "Host: " + host + "\r\n"

    // Добавляем все дополнительные заголовки (если они есть)
    if headers != "" {
        request += headers
    }

    // Пустая строка между заголовками и телом
    request += "\r\n"

    // Добавляем тело (если оно есть)
    if body != "" {
        request += body
    }

    return request
}