задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте функцию MakeCurlCommand(method, url, headers, body string) string,
которая формирует строку команды для утилиты curl на основе переданных параметров:

method — HTTP-метод (например, "GET", "POST"). Для "GET" параметр -X можно не указывать
url — полный адрес (например, "https://example.com/api/data")
headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например, "Content-Type: application/json\nAuthorization: Bearer xyz\n")
body — тело запроса (использовать только если оно не пустое; передавать через --data '...')
Функция должна вернуть строку команды, которую можно выполнить в терминале.
Каждый заголовок должен указываться через -H '...', тело через --data '...',
URL — последний параметр. Заголовки и тело должны корректно экранироваться одинарными кавычками.

Примечания
Пример:

Вызов:

MakeCurlCommand(
  "POST",
  "https://example.com/api/users",
  "Content-Type: application/json\nAuthorization: Bearer abc123\n",
  `{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`
)
Должен вернуть строку:

curl -X POST -H 'Content-Type: application/json' -H 'Authorization: Bearer abc123' --data '{"name":"John Doe","email":"johndoe@example.com","password":"123456"}' https://example.com/api/users