задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте функцию DecodeStudentFromReader(r io.Reader) (Student, error),
которая принимает источник данных, реализующий интерфейс io.Reader, и читает из него JSON-строку, декодируя её в структуру Student:

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}
Если данные некорректны, функция возвращает ошибку.