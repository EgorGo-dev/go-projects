задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте функцию LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64), которая логирует HTTP-запрос с уровнем Info и полями:

"method" — HTTP-метод (GET, POST и т.д.)
"path" — путь запроса (/api/user)
"status" — HTTP-статус (например, 200)
"duration_ms" — длительность обработки запроса в миллисекундах