задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте HTTP-сервер на Go, который:

Обрабатывает запросы по пути /hello и возвращает строку Hello, middleware!.
Использует middleware-функцию Logger(next http.Handler) http.Handler, которая логирует сообщение incoming request, HTTP-метод и путь каждого запроса через пакет slog перед передачей управления следующему обработчику.
Пример работы:

curl -X POST http://localhost:8080/hello
В логах сервера при таком запросе должна появиться строка вида:

level=INFO msg="incoming request" method=POST path=/hello