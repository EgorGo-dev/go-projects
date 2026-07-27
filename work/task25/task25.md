задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте функцию BuildHTTPRequest(method, url, host, headers, body string) string,
которая собирает строку простого HTTP-запроса из переданных параметров:

method — HTTP-метод (например, "GET", "POST")
url — путь к ресурсу (например, "/api/data")
host — заголовок Host (например, "example.com")
headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например, "Content-Type: application/json\r\nAuthorization: Bearer xyz\r\n")
body — тело запроса (может быть пустым).
Функция должна вернуть корректную строку HTTP-запроса, в которой:

первая строка — "<METHOD> <URL> HTTP/1.1"
далее строка "Host: <host>"
далее идут дополнительные заголовки (если есть)
затем пустая строка (разделяет заголовки и тело)
и далее тело запроса (если оно есть)
Примечания
Пример:

Вызов:

BuildHTTPRequest(
  "POST",
  "/api/users",
  "example.com",
  "Content-Type: application/json\r\nAuthorization: Bearer abc123\r\n",
  `{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`
)
Должен вернуть строку:

POST /api/users HTTP/1.1\r\n
Host: example.com\r\n
Content-Type: application/json\r\n
Authorization: Bearer abc123\r\n
\r\n
{"name":"John Doe","email":"johndoe@example.com","password":"123456"}