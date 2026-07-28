задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте HTTP-сервер с помощью http.HandleFunc, который по запросу на адрес /echo возвращает значение параметра msg из строки запроса.
Если параметр msg отсутствует или пустой — вернуть текст "empty".

Примеры:

curl "http://localhost:8080/echo?msg=hello"
# hello

curl "http://localhost:8080/echo?msg="
# empty

curl "http://localhost:8080/echo"
# empty